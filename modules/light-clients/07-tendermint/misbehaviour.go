package tendermint

import (
	"time"

	errorsmod "cosmossdk.io/errors"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"

	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v10/modules/core/24-host"
	"github.com/cosmos/ibc-go/v10/modules/core/exported"
)

var _ exported.ClientMessage = (*Misbehaviour)(nil)

// FrozenHeight is same for all misbehaviour
var FrozenHeight = clienttypes.NewHeight(0, 1)

// NewMisbehaviour creates a new Misbehaviour instance.
func NewMisbehaviour(clientID string, header1, header2 *Header) *Misbehaviour {
	return &Misbehaviour{
		ClientId: clientID,
		Header1:  header1,
		Header2:  header2,
	}
}

// ClientType is Tendermint light client
func (Misbehaviour) ClientType() string {
	return exported.Tendermint
}

// GetTime returns the timestamp at which misbehaviour occurred. It uses the
// maximum value from both headers to prevent producing an invalid header outside
// of the misbehaviour age range.
func (m Misbehaviour) GetTime() time.Time {
	t1, t2 := m.Header1.GetTime(), m.Header2.GetTime()
	if t1.After(t2) {
		return t1
	}
	return t2
}

// ValidateBasic implements Misbehaviour interface
func (m Misbehaviour) ValidateBasic() error {
	if m.Header1 == nil {
		return errorsmod.Wrap(ErrInvalidHeader, "misbehaviour Header1 cannot be nil")
	}
	if m.Header2 == nil {
		return errorsmod.Wrap(ErrInvalidHeader, "misbehaviour Header2 cannot be nil")
	}
	if m.Header1.TrustedHeight.RevisionHeight == 0 {
		return errorsmod.Wrapf(ErrInvalidHeaderHeight, "misbehaviour Header1 cannot have zero revision height")
	}
	if m.Header2.TrustedHeight.RevisionHeight == 0 {
		return errorsmod.Wrapf(ErrInvalidHeaderHeight, "misbehaviour Header2 cannot have zero revision height")
	}
	if m.Header1.TrustedValidators == nil {
		return errorsmod.Wrap(ErrInvalidValidatorSet, "trusted validator set in Header1 cannot be empty")
	}
	if m.Header2.TrustedValidators == nil {
		return errorsmod.Wrap(ErrInvalidValidatorSet, "trusted validator set in Header2 cannot be empty")
	}
	if m.Header1.Header.ChainID != m.Header2.Header.ChainID {
		return errorsmod.Wrap(clienttypes.ErrInvalidMisbehaviour, "headers must have identical chainIDs")
	}

	if err := host.ClientIdentifierValidator(m.ClientId); err != nil {
		return errorsmod.Wrap(err, "misbehaviour client ID is invalid")
	}

	// ValidateBasic on both validators
	if err := m.Header1.ValidateBasic(); err != nil {
		return errorsmod.Wrap(
			clienttypes.ErrInvalidMisbehaviour,
			errorsmod.Wrap(err, "header 1 failed validation").Error(),
		)
	}
	if err := m.Header2.ValidateBasic(); err != nil {
		return errorsmod.Wrap(
			clienttypes.ErrInvalidMisbehaviour,
			errorsmod.Wrap(err, "header 2 failed validation").Error(),
		)
	}
	// Ensure that Height1 is greater than or equal to Height2
	if m.Header1.GetHeight().LT(m.Header2.GetHeight()) {
		return errorsmod.Wrapf(clienttypes.ErrInvalidMisbehaviour, "Header1 height is less than Header2 height (%s < %s)", m.Header1.GetHeight(), m.Header2.GetHeight())
	}

	blockID1, err := cmttypes.BlockIDFromProto(&m.Header1.Commit.BlockID)
	if err != nil {
		return errorsmod.Wrap(err, "invalid block ID from header 1 in misbehaviour")
	}
	blockID2, err := cmttypes.BlockIDFromProto(&m.Header2.Commit.BlockID)
	if err != nil {
		return errorsmod.Wrap(err, "invalid block ID from header 2 in misbehaviour")
	}

	if err := validCommit(m.Header1.Header.ChainID, *blockID1,
		m.Header1.Commit, m.Header1.ValidatorSet); err != nil {
		return err
	}
	return validCommit(m.Header2.Header.ChainID, *blockID2,
		m.Header2.Commit, m.Header2.ValidatorSet)
}

// validCommit checks if the given commit is a valid commit from the passed-in validatorset
func validCommit(chainID string, blockID cmttypes.BlockID, commit *cmtproto.Commit, valSet *cmtproto.ValidatorSet) error {
	tmCommit, err := cmttypes.CommitFromProto(commit)
	if err != nil {
		return errorsmod.Wrap(err, "commit is not tendermint commit type")
	}
	tmValset, err := cmttypes.ValidatorSetFromProto(valSet)
	if err != nil {
		return errorsmod.Wrap(err, "validator set is not tendermint validator set type")
	}

	if err := tmValset.VerifyCommitLight(chainID, blockID, tmCommit.Height, tmCommit); err != nil {
		return errorsmod.Wrap(clienttypes.ErrInvalidMisbehaviour, "validator set did not commit to header")
	}

	return nil
}

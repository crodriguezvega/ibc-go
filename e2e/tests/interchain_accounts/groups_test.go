//go:build !test_e2e

package interchainaccounts

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	interchaintest "github.com/cosmos/interchaintest/v10"
	"github.com/cosmos/interchaintest/v10/ibc"
	test "github.com/cosmos/interchaintest/v10/testutil"
	testifysuite "github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	grouptypes "github.com/cosmos/cosmos-sdk/x/group"

	"github.com/cosmos/ibc-go/e2e/testsuite"
	"github.com/cosmos/ibc-go/e2e/testsuite/query"
	"github.com/cosmos/ibc-go/e2e/testvalues"
	controllertypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v10/testing"
)

const (
	// DefaultGroupMemberWeight is the members voting weight.
	// A group members weight is used in the sum of `YES` votes required to meet a decision policy threshold.
	DefaultGroupMemberWeight = "1"

	// DefaultGroupThreshold is the minimum weighted sum of `YES` votes that must be met or
	// exceeded for a proposal to succeed.
	DefaultGroupThreshold = "1"

	// DefaultMetadata defines a reusable metadata string for testing purposes
	DefaultMetadata = "custom metadata"

	// DefaultMinExecutionPeriod is the minimum duration after the proposal submission
	// where members can start sending MsgExec. This means that the window for
	// sending a MsgExec transaction is:
	// `[ submission + min_execution_period ; submission + voting_period + max_execution_period]`
	// where max_execution_period is an app-specific config, defined in the keeper.
	// If not set, min_execution_period will default to 0.
	DefaultMinExecutionPeriod = time.Duration(0)

	// DefaultVotingPeriod is the duration from submission of a proposal to the end of voting period
	// Within this times votes can be submitted with MsgVote.
	DefaultVotingPeriod = time.Minute

	// InitialGroupID is the first group ID generated by x/group
	InitialGroupID = 1

	// InitialProposalID is the first group proposal ID generated by x/group
	InitialProposalID = 1
)

// compatibility:from_version: v7.10.0
func TestInterchainAccountsGroupsTestSuite(t *testing.T) {
	testifysuite.Run(t, new(InterchainAccountsGroupsTestSuite))
}

type InterchainAccountsGroupsTestSuite struct {
	testsuite.E2ETestSuite
}

// SetupSuite sets up chains for the current test suite
func (s *InterchainAccountsGroupsTestSuite) SetupSuite() {
	s.SetupChains(context.TODO(), 2, nil)
}

func (s *InterchainAccountsGroupsTestSuite) QueryGroupPolicyAddress(ctx context.Context, chain ibc.Chain) string {
	res, err := query.GRPCQuery[grouptypes.QueryGroupPoliciesByGroupResponse](ctx, chain, &grouptypes.QueryGroupPoliciesByGroupRequest{
		GroupId: InitialGroupID, // always use the initial group id
	})
	s.Require().NoError(err)

	return res.GroupPolicies[0].Address
}

func (s *InterchainAccountsGroupsTestSuite) TestInterchainAccountsGroupsIntegration() {
	t := s.T()
	ctx := context.TODO()

	var (
		groupPolicyAddr   string
		interchainAccAddr string
		err               error
	)

	testName := t.Name()
	s.CreatePaths(ibc.DefaultClientOpts(), s.TransferChannelOptions(), testName)
	relayer := s.GetRelayerForTest(testName)

	chainA, chainB := s.GetChains()

	chainAWallet := s.CreateUserOnChainA(ctx, testvalues.StartingTokenAmount)
	chainAAddress := chainAWallet.FormattedAddress()

	chainBWallet := s.CreateUserOnChainB(ctx, testvalues.StartingTokenAmount)
	chainBAddress := chainBWallet.FormattedAddress()

	t.Run("create group with new threshold decision policy", func(t *testing.T) {
		members := []grouptypes.MemberRequest{
			{
				Address: chainAAddress,
				Weight:  DefaultGroupMemberWeight,
			},
		}

		decisionPolicy := grouptypes.NewThresholdDecisionPolicy(DefaultGroupThreshold, DefaultVotingPeriod, DefaultMinExecutionPeriod)
		msgCreateGroupWithPolicy, err := grouptypes.NewMsgCreateGroupWithPolicy(chainAAddress, members, DefaultMetadata, DefaultMetadata, true, decisionPolicy)
		s.Require().NoError(err)

		txResp := s.BroadcastMessages(ctx, chainA, chainAWallet, msgCreateGroupWithPolicy)
		s.AssertTxSuccess(txResp)
	})

	t.Run("submit proposal for MsgRegisterInterchainAccount", func(t *testing.T) {
		groupPolicyAddr = s.QueryGroupPolicyAddress(ctx, chainA)
		msgRegisterAccount := controllertypes.NewMsgRegisterInterchainAccount(ibctesting.FirstConnectionID, groupPolicyAddr, icatypes.NewDefaultMetadataString(ibctesting.FirstConnectionID, ibctesting.FirstConnectionID), channeltypes.ORDERED)

		msgSubmitProposal, err := grouptypes.NewMsgSubmitProposal(groupPolicyAddr, []string{chainAAddress}, []sdk.Msg{msgRegisterAccount}, DefaultMetadata, grouptypes.Exec_EXEC_UNSPECIFIED, "e2e groups proposal: for MsgRegisterInterchainAccount", "e2e groups proposal: for MsgRegisterInterchainAccount")
		s.Require().NoError(err)

		txResp := s.BroadcastMessages(ctx, chainA, chainAWallet, msgSubmitProposal)
		s.AssertTxSuccess(txResp)
	})

	t.Run("vote and exec proposal", func(t *testing.T) {
		msgVote := &grouptypes.MsgVote{
			ProposalId: InitialProposalID,
			Voter:      chainAAddress,
			Option:     grouptypes.VOTE_OPTION_YES,
			Exec:       grouptypes.Exec_EXEC_TRY,
		}

		txResp := s.BroadcastMessages(ctx, chainA, chainAWallet, msgVote)
		s.AssertTxSuccess(txResp)
	})

	t.Run("start relayer", func(t *testing.T) {
		s.StartRelayer(relayer, testName)
	})

	t.Run("verify interchain account registration success", func(t *testing.T) {
		interchainAccAddr, err = query.InterchainAccount(ctx, chainA, groupPolicyAddr, ibctesting.FirstConnectionID)
		s.Require().NotEmpty(interchainAccAddr)
		s.Require().NoError(err)

		channels, err := relayer.GetChannels(ctx, s.GetRelayerExecReporter(), chainA.Config().ChainID)
		s.Require().NoError(err)
		s.Require().Len(channels, 2) // 1 transfer (created by default), 1 interchain-accounts
	})

	t.Run("fund interchain account wallet", func(t *testing.T) {
		err := chainB.SendFunds(ctx, interchaintest.FaucetAccountKeyName, ibc.WalletAmount{
			Address: interchainAccAddr,
			Amount:  sdkmath.NewInt(testvalues.StartingTokenAmount),
			Denom:   chainB.Config().Denom,
		})
		s.Require().NoError(err)
	})

	t.Run("submit proposal for MsgSendTx", func(t *testing.T) {
		msgBankSend := &banktypes.MsgSend{
			FromAddress: interchainAccAddr,
			ToAddress:   chainBAddress,
			Amount:      sdk.NewCoins(testvalues.DefaultTransferAmount(chainB.Config().Denom)),
		}

		cdc := testsuite.Codec()

		bz, err := icatypes.SerializeCosmosTx(cdc, []proto.Message{msgBankSend}, icatypes.EncodingProtobuf)
		s.Require().NoError(err)

		packetData := icatypes.InterchainAccountPacketData{
			Type: icatypes.EXECUTE_TX,
			Data: bz,
			Memo: "e2e",
		}

		msgSubmitTx := controllertypes.NewMsgSendTx(groupPolicyAddr, ibctesting.FirstConnectionID, uint64(time.Hour.Nanoseconds()), packetData)
		msgSubmitProposal, err := grouptypes.NewMsgSubmitProposal(groupPolicyAddr, []string{chainAAddress}, []sdk.Msg{msgSubmitTx}, DefaultMetadata, grouptypes.Exec_EXEC_UNSPECIFIED, "e2e groups proposal: for MsgRegisterInterchainAccount", "e2e groups proposal: for MsgRegisterInterchainAccount")
		s.Require().NoError(err)

		txResp := s.BroadcastMessages(ctx, chainA, chainAWallet, msgSubmitProposal)
		s.AssertTxSuccess(txResp)
	})

	t.Run("vote and exec proposal", func(t *testing.T) {
		msgVote := &grouptypes.MsgVote{
			ProposalId: InitialProposalID + 1,
			Voter:      chainAAddress,
			Option:     grouptypes.VOTE_OPTION_YES,
			Exec:       grouptypes.Exec_EXEC_TRY,
		}

		txResp := s.BroadcastMessages(ctx, chainA, chainAWallet, msgVote)
		s.AssertTxSuccess(txResp)
	})

	t.Run("verify tokens transferred", func(t *testing.T) {
		s.Require().NoError(test.WaitForBlocks(ctx, 10, chainA, chainB), "failed to wait for blocks")
		balance, err := query.Balance(ctx, chainB, chainBAddress, chainB.Config().Denom)

		s.Require().NoError(err)

		expected := testvalues.IBCTransferAmount + testvalues.StartingTokenAmount
		s.Require().Equal(expected, balance.Int64())

		balance, err = query.Balance(ctx, chainB, interchainAccAddr, chainB.Config().Denom)
		s.Require().NoError(err)

		expected = testvalues.StartingTokenAmount - testvalues.IBCTransferAmount
		s.Require().Equal(expected, balance.Int64())
	})
}

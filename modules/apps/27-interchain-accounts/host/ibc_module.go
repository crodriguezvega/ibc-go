package host

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/host/keeper"
	"github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"
	ibcerrors "github.com/cosmos/ibc-go/v10/modules/core/errors"
	ibcexported "github.com/cosmos/ibc-go/v10/modules/core/exported"
)

var (
	_ porttypes.IBCModule             = (*IBCModule)(nil)
	_ porttypes.PacketDataUnmarshaler = (*IBCModule)(nil)
)

// IBCModule implements the ICS26 interface for interchain accounts host chains
type IBCModule struct {
	keeper *keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the associated keeper
func NewIBCModule(k *keeper.Keeper) *IBCModule {
	return &IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (*IBCModule) OnChanOpenInit(
	_ sdk.Context,
	_ channeltypes.Order,
	_ []string,
	_ string,
	_ string,
	_ channeltypes.Counterparty,
	_ string,
) (string, error) {
	return "", errorsmod.Wrap(icatypes.ErrInvalidChannelFlow, "channel handshake must be initiated by controller chain")
}

// OnChanOpenTry implements the IBCModule interface
func (im *IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	if !im.keeper.GetParams(ctx).HostEnabled {
		return "", types.ErrHostSubModuleDisabled
	}

	return im.keeper.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCModule interface
func (*IBCModule) OnChanOpenAck(
	_ sdk.Context,
	_,
	_ string,
	_ string,
	_ string,
) error {
	return errorsmod.Wrap(icatypes.ErrInvalidChannelFlow, "channel handshake must be initiated by controller chain")
}

// OnChanOpenConfirm implements the IBCModule interface
func (im *IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	if !im.keeper.GetParams(ctx).HostEnabled {
		return types.ErrHostSubModuleDisabled
	}

	return im.keeper.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCModule interface
func (*IBCModule) OnChanCloseInit(
	_ sdk.Context,
	_ string,
	_ string,
) error {
	// Disallow user-initiated channel closing for interchain account channels
	return errorsmod.Wrap(ibcerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (im *IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.keeper.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCModule interface
func (im *IBCModule) OnRecvPacket(
	ctx sdk.Context,
	_ string,
	packet channeltypes.Packet,
	_ sdk.AccAddress,
) ibcexported.Acknowledgement {
	if !im.keeper.GetParams(ctx).HostEnabled {
		im.keeper.Logger(ctx).Info("host submodule is disabled")
		keeper.EmitHostDisabledEvent(ctx, packet)
		return channeltypes.NewErrorAcknowledgement(types.ErrHostSubModuleDisabled)
	}

	txResponse, err := im.keeper.OnRecvPacket(ctx, packet)
	ack := channeltypes.NewResultAcknowledgement(txResponse)
	if err != nil {
		ack = channeltypes.NewErrorAcknowledgement(err)
		im.keeper.Logger(ctx).Error(fmt.Sprintf("%s sequence %d", err.Error(), packet.Sequence))
	} else {
		im.keeper.Logger(ctx).Info("successfully handled packet", "sequence", packet.Sequence)
	}

	// Emit an event indicating a successful or failed acknowledgement.
	keeper.EmitAcknowledgementEvent(ctx, packet, ack, err)

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface
func (*IBCModule) OnAcknowledgementPacket(
	_ sdk.Context,
	_ string,
	_ channeltypes.Packet,
	_ []byte,
	_ sdk.AccAddress,
) error {
	return errorsmod.Wrap(icatypes.ErrInvalidChannelFlow, "cannot receive acknowledgement on a host channel end, a host chain does not send a packet over the channel")
}

// OnTimeoutPacket implements the IBCModule interface
func (*IBCModule) OnTimeoutPacket(
	_ sdk.Context,
	_ string,
	_ channeltypes.Packet,
	_ sdk.AccAddress,
) error {
	return errorsmod.Wrap(icatypes.ErrInvalidChannelFlow, "cannot cause a packet timeout on a host channel end, a host chain does not send a packet over the channel")
}

// UnmarshalPacketData attempts to unmarshal the provided packet data bytes
// into an InterchainAccountPacketData. This function implements the optional
// PacketDataUnmarshaler interface required for ADR 008 support.
func (im *IBCModule) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	var data icatypes.InterchainAccountPacketData
	err := data.UnmarshalJSON(bz)
	if err != nil {
		return nil, "", err
	}

	version, ok := im.keeper.GetAppVersion(ctx, portID, channelID)
	if !ok {
		return nil, "", errorsmod.Wrapf(ibcerrors.ErrNotFound, "app version not found for port %s and channel %s", portID, channelID)
	}

	return data, version, nil
}

// SetICS4Wrapper sets the ICS4Wrapper for the IBCModule.
func (im *IBCModule) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	if wrapper == nil {
		panic("ICS4Wrapper cannot be nil")
	}
	im.keeper.WithICS4Wrapper(wrapper)
}

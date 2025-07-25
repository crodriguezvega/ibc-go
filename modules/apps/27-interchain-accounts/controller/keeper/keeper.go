package keeper

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	corestore "cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/controller/types"
	genesistypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/genesis/types"
	icatypes "github.com/cosmos/ibc-go/v10/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"
	ibcerrors "github.com/cosmos/ibc-go/v10/modules/core/errors"
	"github.com/cosmos/ibc-go/v10/modules/core/exported"
)

// Keeper defines the IBC interchain accounts controller keeper
type Keeper struct {
	storeService  corestore.KVStoreService
	cdc           codec.Codec
	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper icatypes.ChannelKeeper

	msgRouter icatypes.MessageRouter

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new interchain accounts controller Keeper instance
func NewKeeper(
	cdc codec.Codec, storeService corestore.KVStoreService,
	channelKeeper icatypes.ChannelKeeper,
	msgRouter icatypes.MessageRouter, authority string,
) *Keeper {
	if strings.TrimSpace(authority) == "" {
		panic(errors.New("authority must be non-empty"))
	}

	return &Keeper{
		storeService: storeService,
		cdc:          cdc,
		// Defaults to using the channel keeper as the ICS4Wrapper
		// This can be overridden later with WithICS4Wrapper (e.g. by the middleware stack wiring)
		ics4Wrapper:   channelKeeper,
		channelKeeper: channelKeeper,
		msgRouter:     msgRouter,
		authority:     authority,
	}
}

// WithICS4Wrapper sets the ICS4Wrapper. This function may be used after
// the keepers creation to set the middleware which is above this module
// in the IBC application stack.
func (k *Keeper) WithICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	k.ics4Wrapper = wrapper
}

// GetICS4Wrapper returns the ICS4Wrapper.
func (k *Keeper) GetICS4Wrapper() porttypes.ICS4Wrapper {
	return k.ics4Wrapper
}

// Logger returns the application logger, scoped to the associated module
func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s-%s", exported.ModuleName, icatypes.ModuleName))
}

// GetConnectionID returns the connection id for the given port and channelIDs.
func (k *Keeper) GetConnectionID(ctx sdk.Context, portID, channelID string) (string, error) {
	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
	if !found {
		return "", errorsmod.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", portID, channelID)
	}
	return channel.ConnectionHops[0], nil
}

// GetAllPorts returns all ports to which the interchain accounts controller module is bound. Used in ExportGenesis
func (k *Keeper) GetAllPorts(ctx sdk.Context) []string {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := storetypes.KVStorePrefixIterator(store, []byte(icatypes.PortKeyPrefix))
	defer sdk.LogDeferred(k.Logger(ctx), func() error { return iterator.Close() })

	var ports []string
	for ; iterator.Valid(); iterator.Next() {
		keySplit := strings.Split(string(iterator.Key()), "/")

		ports = append(ports, keySplit[1])
	}

	return ports
}

// setPort sets the provided portID in state
func (k *Keeper) setPort(ctx sdk.Context, portID string) {
	store := k.storeService.OpenKVStore(ctx)
	if err := store.Set(icatypes.KeyPort(portID), []byte{0x01}); err != nil {
		panic(err)
	}
}

// GetAppVersion calls the ICS4Wrapper GetAppVersion function.
func (k *Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return k.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}

// GetActiveChannelID retrieves the active channelID from the store, keyed by the provided connectionID and portID
func (k *Keeper) GetActiveChannelID(ctx sdk.Context, connectionID, portID string) (string, bool) {
	store := k.storeService.OpenKVStore(ctx)
	key := icatypes.KeyActiveChannel(portID, connectionID)

	bz, err := store.Get(key)
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return "", false
	}

	return string(bz), true
}

// GetOpenActiveChannel retrieves the active channelID from the store, keyed by the provided connectionID and portID & checks if the channel in question is in state OPEN
func (k *Keeper) GetOpenActiveChannel(ctx sdk.Context, connectionID, portID string) (string, bool) {
	channelID, found := k.GetActiveChannelID(ctx, connectionID, portID)
	if !found {
		return "", false
	}

	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)

	if found && channel.State == channeltypes.OPEN {
		return channelID, true
	}

	return "", false
}

// IsActiveChannelClosed retrieves the active channel from the store and returns true if the channel state is CLOSED, otherwise false
func (k *Keeper) IsActiveChannelClosed(ctx sdk.Context, connectionID, portID string) bool {
	channelID, found := k.GetActiveChannelID(ctx, connectionID, portID)
	if !found {
		return false
	}

	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
	return found && channel.State == channeltypes.CLOSED
}

// GetAllActiveChannels returns a list of all active interchain accounts controller channels and their associated connection and port identifiers
func (k *Keeper) GetAllActiveChannels(ctx sdk.Context) []genesistypes.ActiveChannel {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := storetypes.KVStorePrefixIterator(store, []byte(icatypes.ActiveChannelKeyPrefix))
	defer sdk.LogDeferred(k.Logger(ctx), func() error { return iterator.Close() })

	var activeChannels []genesistypes.ActiveChannel
	for ; iterator.Valid(); iterator.Next() {
		keySplit := strings.Split(string(iterator.Key()), "/")

		portID := keySplit[1]
		connectionID := keySplit[2]
		channelID := string(iterator.Value())

		ch := genesistypes.ActiveChannel{
			ConnectionId:        connectionID,
			PortId:              portID,
			ChannelId:           channelID,
			IsMiddlewareEnabled: k.IsMiddlewareEnabled(ctx, portID, connectionID),
		}

		activeChannels = append(activeChannels, ch)
	}

	return activeChannels
}

// SetActiveChannelID stores the active channelID, keyed by the provided connectionID and portID
func (k *Keeper) SetActiveChannelID(ctx sdk.Context, connectionID, portID, channelID string) {
	store := k.storeService.OpenKVStore(ctx)
	if err := store.Set(icatypes.KeyActiveChannel(portID, connectionID), []byte(channelID)); err != nil {
		panic(err)
	}
}

// IsActiveChannel returns true if there exists an active channel for the provided connectionID and portID, otherwise false
func (k *Keeper) IsActiveChannel(ctx sdk.Context, connectionID, portID string) bool {
	_, ok := k.GetActiveChannelID(ctx, connectionID, portID)
	return ok
}

// GetInterchainAccountAddress retrieves the InterchainAccount address from the store associated with the provided connectionID and portID
func (k *Keeper) GetInterchainAccountAddress(ctx sdk.Context, connectionID, portID string) (string, bool) {
	store := k.storeService.OpenKVStore(ctx)
	key := icatypes.KeyOwnerAccount(portID, connectionID)

	bz, err := store.Get(key)
	if err != nil {
		panic(err)
	}
	if len(bz) == 0 {
		return "", false
	}

	return string(bz), true
}

// GetAllInterchainAccounts returns a list of all registered interchain account addresses and their associated connection and controller port identifiers
func (k *Keeper) GetAllInterchainAccounts(ctx sdk.Context) []genesistypes.RegisteredInterchainAccount {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	iterator := storetypes.KVStorePrefixIterator(store, []byte(icatypes.OwnerKeyPrefix))

	var interchainAccounts []genesistypes.RegisteredInterchainAccount
	for ; iterator.Valid(); iterator.Next() {
		keySplit := strings.Split(string(iterator.Key()), "/")

		acc := genesistypes.RegisteredInterchainAccount{
			ConnectionId:   keySplit[2],
			PortId:         keySplit[1],
			AccountAddress: string(iterator.Value()),
		}

		interchainAccounts = append(interchainAccounts, acc)
	}

	return interchainAccounts
}

// SetInterchainAccountAddress stores the InterchainAccount address, keyed by the associated connectionID and portID
func (k *Keeper) SetInterchainAccountAddress(ctx sdk.Context, connectionID, portID, address string) {
	store := k.storeService.OpenKVStore(ctx)
	if err := store.Set(icatypes.KeyOwnerAccount(portID, connectionID), []byte(address)); err != nil {
		panic(err)
	}
}

// IsMiddlewareEnabled returns true if the underlying application callbacks are enabled for given port and connection identifier pair, otherwise false
func (k *Keeper) IsMiddlewareEnabled(ctx sdk.Context, portID, connectionID string) bool {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(icatypes.KeyIsMiddlewareEnabled(portID, connectionID))
	if err != nil {
		panic(err)
	}
	return bytes.Equal(icatypes.MiddlewareEnabled, bz)
}

// IsMiddlewareDisabled returns true if the underlying application callbacks are disabled for the given port and connection identifier pair, otherwise false
func (k *Keeper) IsMiddlewareDisabled(ctx sdk.Context, portID, connectionID string) bool {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(icatypes.KeyIsMiddlewareEnabled(portID, connectionID))
	if err != nil {
		panic(err)
	}
	return bytes.Equal(icatypes.MiddlewareDisabled, bz)
}

// SetMiddlewareEnabled stores a flag to indicate that the underlying application callbacks should be enabled for the given port and connection identifier pair
func (k *Keeper) SetMiddlewareEnabled(ctx sdk.Context, portID, connectionID string) {
	store := k.storeService.OpenKVStore(ctx)
	if err := store.Set(icatypes.KeyIsMiddlewareEnabled(portID, connectionID), icatypes.MiddlewareEnabled); err != nil {
		panic(err)
	}
}

// SetMiddlewareDisabled stores a flag to indicate that the underlying application callbacks should be disabled for the given port and connection identifier pair
func (k *Keeper) SetMiddlewareDisabled(ctx sdk.Context, portID, connectionID string) {
	store := k.storeService.OpenKVStore(ctx)
	if err := store.Set(icatypes.KeyIsMiddlewareEnabled(portID, connectionID), icatypes.MiddlewareDisabled); err != nil {
		panic(err)
	}
}

// DeleteMiddlewareEnabled deletes the middleware enabled flag stored in state
func (k *Keeper) DeleteMiddlewareEnabled(ctx sdk.Context, portID, connectionID string) {
	store := k.storeService.OpenKVStore(ctx)
	if err := store.Delete(icatypes.KeyIsMiddlewareEnabled(portID, connectionID)); err != nil {
		panic(err)
	}
}

// GetAuthority returns the ica/controller submodule's authority.
func (k *Keeper) GetAuthority() string {
	return k.authority
}

// getAppMetadata retrieves the interchain accounts channel metadata from the store associated with the provided portID and channelID
func (k *Keeper) getAppMetadata(ctx sdk.Context, portID, channelID string) (icatypes.Metadata, error) {
	appVersion, found := k.GetAppVersion(ctx, portID, channelID)
	if !found {
		return icatypes.Metadata{}, errorsmod.Wrapf(ibcerrors.ErrNotFound, "app version not found for port %s and channel %s", portID, channelID)
	}

	return icatypes.MetadataFromVersion(appVersion)
}

// GetParams returns the current ica/controller submodule parameters.
func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get([]byte(types.ParamsKey))
	if err != nil {
		panic(err)
	}
	if bz == nil { // only panic on unset params and not on empty params
		panic(errors.New("ica/controller params are not set in store"))
	}

	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams sets the ica/controller submodule parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := k.storeService.OpenKVStore(ctx)
	bz := k.cdc.MustMarshal(&params)
	if err := store.Set([]byte(types.ParamsKey), bz); err != nil {
		panic(err)
	}
}

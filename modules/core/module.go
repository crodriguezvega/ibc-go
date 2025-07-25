package ibc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	ibcclient "github.com/cosmos/ibc-go/v10/modules/core/02-client"
	clientkeeper "github.com/cosmos/ibc-go/v10/modules/core/02-client/keeper"
	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	clientv2keeper "github.com/cosmos/ibc-go/v10/modules/core/02-client/v2/keeper"
	clientv2types "github.com/cosmos/ibc-go/v10/modules/core/02-client/v2/types"
	connectionkeeper "github.com/cosmos/ibc-go/v10/modules/core/03-connection/keeper"
	connectiontypes "github.com/cosmos/ibc-go/v10/modules/core/03-connection/types"
	channelkeeper "github.com/cosmos/ibc-go/v10/modules/core/04-channel/keeper"
	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	channelkeeperv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/keeper"
	channeltypesv2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v10/modules/core/client/cli"
	"github.com/cosmos/ibc-go/v10/modules/core/exported"
	"github.com/cosmos/ibc-go/v10/modules/core/keeper"
	"github.com/cosmos/ibc-go/v10/modules/core/simulation"
	"github.com/cosmos/ibc-go/v10/modules/core/types"
)

var (
	_ module.AppModule           = (*AppModule)(nil)
	_ module.AppModuleBasic      = (*AppModuleBasic)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasName             = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)
	_ module.HasServices         = (*AppModule)(nil)
	_ module.HasProposalMsgs     = (*AppModule)(nil)
	_ appmodule.AppModule        = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker  = (*AppModule)(nil)
)

// AppModuleBasic defines the basic application module used by the ibc module.
type AppModuleBasic struct{}

// Name returns the ibc module's name.
func (AppModuleBasic) Name() string {
	return exported.ModuleName
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (AppModule) IsAppModule() {}

// RegisterLegacyAminoCodec implements AppModuleBasic interface.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	clienttypes.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the ibc
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the ibc module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var gs types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &gs); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", exported.ModuleName, err)
	}

	return gs.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the ibc module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := clienttypes.RegisterQueryHandlerClient(context.Background(), mux, clienttypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = connectiontypes.RegisterQueryHandlerClient(context.Background(), mux, connectiontypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = channeltypes.RegisterQueryHandlerClient(context.Background(), mux, channeltypes.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
	err = channeltypesv2.RegisterQueryHandlerClient(context.Background(), mux, channeltypesv2.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// GetTxCmd returns the root tx command for the ibc module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns no root query command for the ibc module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// AppModule implements an application module for the ibc module.
type AppModule struct {
	AppModuleBasic
	keeper *keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k *keeper.Keeper) AppModule {
	return AppModule{
		keeper: k,
	}
}

// Name returns the ibc module's name.
func (AppModule) Name() string {
	return exported.ModuleName
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	clienttypes.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	clientv2types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	connectiontypes.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	channeltypes.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	channeltypesv2.RegisterMsgServer(cfg.MsgServer(), am.keeper.ChannelKeeperV2)

	clienttypes.RegisterQueryServer(cfg.QueryServer(), clientkeeper.NewQueryServer(am.keeper.ClientKeeper))
	clientv2types.RegisterQueryServer(cfg.QueryServer(), clientv2keeper.NewQueryServer(am.keeper.ClientV2Keeper))
	connectiontypes.RegisterQueryServer(cfg.QueryServer(), connectionkeeper.NewQueryServer(am.keeper.ConnectionKeeper))
	channeltypes.RegisterQueryServer(cfg.QueryServer(), channelkeeper.NewQueryServer(am.keeper.ChannelKeeper))
	channeltypesv2.RegisterQueryServer(cfg.QueryServer(), channelkeeperv2.NewQueryServer(am.keeper.ChannelKeeperV2))

	// This upgrade used to just add default params, since we have deleted it (in consensus version 8 - ibc-go v10),
	// we just return directly to increment the ConsensusVersion as expected
	if err := cfg.RegisterMigration(exported.ModuleName, 5, func(_ sdk.Context) error {
		return nil
	}); err != nil {
		panic(err)
	}

	clientMigrator := clientkeeper.NewMigrator(am.keeper.ClientKeeper)
	if err := cfg.RegisterMigration(exported.ModuleName, 6, clientMigrator.MigrateToStatelessLocalhost); err != nil {
		panic(err)
	}

	channelMigrator := channelkeeper.NewMigrator(am.keeper.ChannelKeeper)
	if err := cfg.RegisterMigration(exported.ModuleName, 7, channelMigrator.Migrate7To8); err != nil {
		panic(err)
	}
}

// InitGenesis performs genesis initialization for the ibc module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
	var gs types.GenesisState
	err := cdc.UnmarshalJSON(bz, &gs)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal %s genesis state: %w", exported.ModuleName, err))
	}
	InitGenesis(ctx, *am.keeper, &gs)
}

// ExportGenesis returns the exported genesis state as raw bytes for the ibc
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(ExportGenesis(ctx, *am.keeper))
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 8 }

// BeginBlock returns the begin blocker for the ibc module.
func (am AppModule) BeginBlock(goCtx context.Context) error {
	ibcclient.BeginBlocker(sdk.UnwrapSDKContext(goCtx), am.keeper.ClientKeeper)
	return nil
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the ibc module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return simulation.ProposalMsgs()
}

// RegisterStoreDecoder registers a decoder for ibc module's types
func (am AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	sdr[exported.StoreKey] = simulation.NewDecodeStore(*am.keeper)
}

// WeightedOperations returns the all the ibc module operations with their respective weights.
func (AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	return nil
}

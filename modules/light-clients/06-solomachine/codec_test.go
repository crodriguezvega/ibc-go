package solomachine_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	solomachine "github.com/cosmos/ibc-go/v9/modules/light-clients/06-solomachine"
)

func TestCodecTypeRegistration(t *testing.T) {
	testCases := []struct {
		name    string
		typeURL string
		expErr  error
	}{
		{
			"success: ClientState",
			sdk.MsgTypeURL(&solomachine.ClientState{}),
			nil,
		},
		{
			"success: ConsensusState",
			sdk.MsgTypeURL(&solomachine.ConsensusState{}),
			nil,
		},
		{
			"success: Header",
			sdk.MsgTypeURL(&solomachine.Header{}),
			nil,
		},
		{
			"success: Misbehaviour",
			sdk.MsgTypeURL(&solomachine.Misbehaviour{}),
			nil,
		},
		{
			"type not registered on codec",
			"ibc.invalid.MsgTypeURL",
			errors.New("unable to resolve type URL ibc.invalid.MsgTypeURL"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			encodingCfg := moduletestutil.MakeTestEncodingConfig(testutil.CodecOptions{}, solomachine.AppModuleBasic{})
			msg, err := encodingCfg.Codec.InterfaceRegistry().Resolve(tc.typeURL)

			if tc.expErr == nil {
				require.NotNil(t, msg)
				require.NoError(t, err)
			} else {
				require.Nil(t, msg)
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErr.Error())
			}
		})
	}
}

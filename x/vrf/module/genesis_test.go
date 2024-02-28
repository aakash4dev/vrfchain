package vrf_test

import (
	"testing"

	keepertest "github.com/aakash4dev/vrfchain/testutil/keeper"
	"github.com/aakash4dev/vrfchain/testutil/nullify"
	"github.com/aakash4dev/vrfchain/x/vrf/module"
	"github.com/aakash4dev/vrfchain/x/vrf/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VrfKeeper(t)
	vrf.InitGenesis(ctx, k, genesisState)
	got := vrf.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

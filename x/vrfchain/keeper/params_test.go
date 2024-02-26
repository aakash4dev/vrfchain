package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/aakash4dev/vrfchain/testutil/keeper"
	"github.com/aakash4dev/vrfchain/x/vrfchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VrfchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

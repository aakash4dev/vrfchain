package keeper

import (
	"github.com/aakash4dev/vrfchain/x/vrf/types"
)

var _ types.QueryServer = Keeper{}

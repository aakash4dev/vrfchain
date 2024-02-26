package keeper

import (
	"github.com/aakash4dev/vrfchain/x/vrfchain/types"
)

var _ types.QueryServer = Keeper{}

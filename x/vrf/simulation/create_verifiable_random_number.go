package simulation

import (
	"math/rand"

	"github.com/aakash4dev/vrfchain/x/vrf/keeper"
	"github.com/aakash4dev/vrfchain/x/vrf/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateVerifiableRandomNumber(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateVerifiableRandomNumber{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateVerifiableRandomNumber simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "CreateVerifiableRandomNumber simulation not implemented"), nil, nil
	}
}

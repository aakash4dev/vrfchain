package vrf

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/aakash4dev/vrfchain/testutil/sample"
	vrfsimulation "github.com/aakash4dev/vrfchain/x/vrf/simulation"
	"github.com/aakash4dev/vrfchain/x/vrf/types"
)

// avoid unused import issue
var (
	_ = vrfsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateVerifiableRandomNumber = "op_weight_msg_create_verifiable_random_number"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVerifiableRandomNumber int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vrfGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vrfGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateVerifiableRandomNumber int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateVerifiableRandomNumber, &weightMsgCreateVerifiableRandomNumber, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVerifiableRandomNumber = defaultWeightMsgCreateVerifiableRandomNumber
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVerifiableRandomNumber,
		vrfsimulation.SimulateMsgCreateVerifiableRandomNumber(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateVerifiableRandomNumber,
			defaultWeightMsgCreateVerifiableRandomNumber,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vrfsimulation.SimulateMsgCreateVerifiableRandomNumber(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

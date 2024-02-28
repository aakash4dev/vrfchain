package vrf

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/aakash4dev/vrfchain/api/vrfchain/vrf"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateVerifiableRandomNumber",
					Use:            "create-verifiable-random-number [shaseed] [publickey] [r] [s] [max-range]",
					Short:          "Send a create-verifiable-random-number tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "shaseed"}, {ProtoField: "publickey"}, {ProtoField: "r"}, {ProtoField: "s"}, {ProtoField: "maxRange"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: voucherv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Balance",
					Use:       "balance [id] [address]",
					Short:     "query balance by id and address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "BalancesByAddress",
					Use:       "balances-by-address [address]",
					Short:     "query all balances by address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "BalancesByVoucher",
					Use:       "balances-by-voucher [id]",
					Short:     "query all balances by voucher",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				{
					RpcMethod: "Voucher",
					Use:       "voucher [id]",
					Short:     "query voucher by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				{
					RpcMethod: "Vouchers",
					Use:       "vouchers",
					Short:     "query all vouchers",
				},
				{
					RpcMethod: "VouchersByIssuer",
					Use:       "vouchers-by-issuer [issuer]",
					Short:     "query vouchers by issuer",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "issuer"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              voucherv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions:    []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "Create",
				//	Use:       "create [metadata]",
				//	Short:     "submit a transaction to create voucher",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "Issue",
				//	Use:       "issue [id] [recipient] [amount] [expiration] [metadata]",
				//	Short:     "submit a transaction to issue vouchers",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//		{ProtoField: "recipient"},
				//		{ProtoField: "amount"},
				//		{ProtoField: "expiration"},
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateIssuer",
				//	Use:       "update-issuer [id] [new-issuer]",
				//	Short:     "submit a transaction to update voucher issuer",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//		{ProtoField: "new_issuer"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateMetadata",
				//	Use:       "update-issuer [id] [new-metadata]",
				//	Short:     "submit a transaction to update voucher metadata",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//		{ProtoField: "new_metadata"},
				//	},
				//},
			},
		},
	}
}

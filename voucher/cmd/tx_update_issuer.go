package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/voucher/types/v1"
)

// TxUpdateIssuerCmd creates and returns the tx update-issuer command.
func TxUpdateIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-issuer [id] [new-issuer]",
		Short: "submit transaction to update voucher issuer",
		Long:  "submit transaction to update voucher issuer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateIssuer{
				Id:        id,
				Issuer:    clientCtx.GetFromAddress().String(),
				NewIssuer: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/authority/types/v1"
)

// TxUpdateCmd creates and returns the tx update-curator command.
func TxUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [new-authority]",
		Short: "submit a transaction to update the authority account",
		Long:  "submit a transaction to update the authority account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdate{
				Authority:    clientCtx.GetFromAddress().String(),
				NewAuthority: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

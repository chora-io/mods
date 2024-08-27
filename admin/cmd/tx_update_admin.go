package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/admin/types/v1"
)

// TxUpdateAdminCmd creates and returns the tx update-admin command.
func TxUpdateAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-admin [new-admin]",
		Short: "submit a transaction to update the admin account",
		Long:  "submit a transaction to update the admin account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateAdmin{
				Admin:    clientCtx.GetFromAddress().String(),
				NewAdmin: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

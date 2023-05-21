package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// TxAddCmd creates and returns the tx add command.
func TxAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [address]",
		Short: "submit a transaction to add a validator",
		Long:  "submit a transaction to add a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgAdd{
				Authority: clientCtx.GetFromAddress().String(),
				Address:   args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

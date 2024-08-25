package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// TxAddValidatorCmd creates and returns the tx add command.
func TxAddValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-validator [address] [metadata]",
		Short: "submit a transaction to add a validator",
		Long:  "submit a transaction to add a validator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgAddValidator{
				Admin:    clientCtx.GetFromAddress().String(),
				Address:  args[0],
				Metadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

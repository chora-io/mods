package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// TxRemoveValidatorCmd creates and returns the tx remove command.
func TxRemoveValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-validator [address]",
		Short: "submit transaction to remove validator",
		Long:  "submit transaction to remove validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgRemoveValidator{
				Operator: clientCtx.GetFromAddress().String(),
				Address:  args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

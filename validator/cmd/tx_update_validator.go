package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// TxUpdateValidatorCmd creates and returns the tx update command.
func TxUpdateValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-validator [new-metadata]",
		Short: "submit a transaction to update validator metadata",
		Long:  "submit a transaction to update validator metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateValidator{
				Address:     clientCtx.GetFromAddress().String(),
				NewMetadata: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

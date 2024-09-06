package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// TxCreateValidatorCmd creates and returns the tx create-validator command.
func TxCreateValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator [metadata]",
		Short: "submit transaction to create validator",
		Long:  "submit transaction to create validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgCreateValidator{
				Operator: clientCtx.GetFromAddress().String(),
				Metadata: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

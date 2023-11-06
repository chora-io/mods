package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/content/types/v1"
)

// TxCreateCmd creates and returns the tx create command.
func TxCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [metadata]",
		Short: "submit a transaction to create content",
		Long:  "submit a transaction to create content",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgCreate{
				Curator:  clientCtx.GetFromAddress().String(),
				Metadata: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

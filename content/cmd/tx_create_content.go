package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/content/types/v1"
)

// TxCreateContentCmd creates and returns the tx create command.
func TxCreateContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-content [hash]",
		Short: "submit transaction to create content",
		Long:  "submit transaction to create content",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgCreateContent{
				Curator: clientCtx.GetFromAddress().String(),
				Hash:    args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

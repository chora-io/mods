package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/content/types/v1"
)

// TxUpdateContentCuratorCmd creates and returns the tx update-curator command.
func TxUpdateContentCuratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-content-curator [hash] [new-curator]",
		Short: "submit transaction to update content curator",
		Long:  "submit transaction to update content curator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateContentCurator{
				Curator:    clientCtx.GetFromAddress().String(),
				Hash:       args[0],
				NewCurator: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

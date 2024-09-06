package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// TxUpdateGovernorCmd creates and returns the tx update command.
func TxUpdateGovernorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-governor-metadata [new-metadata]",
		Short: "submit transaction to update governor metadata",
		Long:  "submit transaction to update governor metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateGovernorMetadata{
				Address:     clientCtx.GetFromAddress().String(),
				NewMetadata: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/agent/types/v1"
)

// TxUpdateAgentMetadataCmd creates and returns the tx update-agent-metadata command.
func TxUpdateAgentMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-agent-metadata [address] [new-metadata]",
		Short: "submit a transaction to update agent metadata",
		Long:  "submit a transaction to update agent metadata",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateAgentMetadata{
				Address:     args[0],
				Admin:       clientCtx.GetFromAddress().String(),
				NewMetadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

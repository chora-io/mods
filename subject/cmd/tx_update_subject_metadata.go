package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// TxUpdateSubjectMetadataCmd creates and returns the tx update-subject-metadata command.
func TxUpdateSubjectMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-subject-metadata [address] [new-metadata]",
		Short: "submit transaction to update subject metadata",
		Long:  "submit transaction to update subject metadata",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateSubjectMetadata{
				Steward:     clientCtx.GetFromAddress().String(),
				Address:     args[0],
				NewMetadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

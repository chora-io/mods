package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// TxUpdateSubjectStewardCmd creates and returns the tx update-subject-steward command.
func TxUpdateSubjectStewardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-subject-steward [address] [new-steward]",
		Short: "submit transaction to update subject steward",
		Long:  "submit transaction to update subject steward",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateSubjectSteward{
				Address:    args[0],
				Steward:    clientCtx.GetFromAddress().String(),
				NewSteward: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

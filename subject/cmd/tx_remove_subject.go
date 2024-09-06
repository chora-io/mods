package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// TxRemoveSubjectCmd creates and returns the tx remove-subject command.
func TxRemoveSubjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-subject [address]",
		Short: "submit transaction to remove subject",
		Long:  "submit transaction to remove subject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgRemoveSubject{
				Steward: clientCtx.GetFromAddress().String(),
				Address: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

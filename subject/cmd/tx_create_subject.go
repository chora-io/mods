package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// TxCreateSubjectCmd creates and returns the tx create-subject command.
func TxCreateSubjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-subject [metadata]",
		Short: "submit transaction to create subject",
		Long:  "submit transaction to create subject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgCreateSubject{
				Steward:  clientCtx.GetFromAddress().String(),
				Metadata: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

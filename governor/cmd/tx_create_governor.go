package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// TxCreateGovernorCmd creates and returns the tx add command.
func TxCreateGovernorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-governor [metadata]",
		Short: "submit a transaction to add a governor",
		Long:  "submit a transaction to add a governor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgCreateGovernor{
				Address:  clientCtx.GetFromAddress().String(),
				Metadata: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

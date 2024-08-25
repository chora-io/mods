package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// TxAddGovernorCmd creates and returns the tx add command.
func TxAddGovernorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-governor [address] [metadata]",
		Short: "submit a transaction to add a governor",
		Long:  "submit a transaction to add a governor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgAddGovernor{
				Admin:    clientCtx.GetFromAddress().String(),
				Address:  args[0],
				Metadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

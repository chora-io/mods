package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// TxRemoveGovernorCmd creates and returns the tx remove command.
func TxRemoveGovernorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-governor [address]",
		Short: "submit a transaction to remove a governor",
		Long:  "submit a transaction to remove a governor",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgRemoveGovernor{
				Admin:   clientCtx.GetFromAddress().String(),
				Address: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

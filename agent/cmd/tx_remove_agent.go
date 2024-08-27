package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/agent/types/v1"
)

// TxRemoveAgentCmd creates and returns the tx remove-agent command.
func TxRemoveAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-agent [address]",
		Short: "submit a transaction to remove an agent",
		Long:  "submit a transaction to remove an agent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgRemoveAgent{
				Address: args[0],
				Admin:   clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/agent/types/v1"
)

// TxUpdateAgentAdminCmd creates and returns the tx update-agent-admin command.
func TxUpdateAgentAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-agent-admin [address] [new-admin]",
		Short: "submit transaction to update agent admin",
		Long:  "submit transaction to update agent admin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v1.MsgUpdateAgentAdmin{
				Address:  args[0],
				Admin:    clientCtx.GetFromAddress().String(),
				NewAdmin: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

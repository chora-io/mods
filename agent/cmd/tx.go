package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/chora-io/mods/agent"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        agent.ModuleName,
		Short:                      "tx commands for the agent module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxCreateAgentCmd(),
		TxUpdateAgentAdminCmd(),
		TxUpdateAgentMetadataCmd(),
	)

	return cmd
}

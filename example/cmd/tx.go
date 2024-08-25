package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/chora-io/mods/example"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        example.ModuleName,
		Short:                      "tx commands for the example module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxCreateCmd(),
		TxDeleteCmd(),
		TxUpdateCuratorCmd(),
		TxUpdateMetadataCmd(),
	)

	return cmd
}

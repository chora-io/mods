package cmd

import (
	"github.com/choraio/mods/content"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        content.ModuleName,
		Short:                      "tx commands for the content module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxCreateCmd(),
		TxUpdateCuratorCmd(),
		TxUpdateMetadataCmd(),
		TxDeleteCmd(),
	)

	return cmd
}

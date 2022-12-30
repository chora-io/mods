package cmd

import (
	"github.com/choraio/mods/content"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// QueryCmd creates and returns the query command.
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        content.ModuleName,
		Short:                      "query commands for the content module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryContentCmd(),
		QueryContentsCmd(),
		QueryContentsByCuratorCmd(),
	)

	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/choraio/mods/example"
)

// QueryCmd creates and returns the query command.
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        example.ModuleName,
		Short:                      "query commands for the example module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryContentCmd(),
		QueryContentByCreatorCmd(),
	)

	return cmd
}

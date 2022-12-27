package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/choraio/mods/geonode"
)

// QueryCmd creates and returns the query command.
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        geonode.ModuleName,
		Short:                      "query commands for the geonode module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryNodeCmd(),
		QueryNodeByCuratorCmd(),
	)

	return cmd
}

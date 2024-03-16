package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/chora-io/mods/geonode"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        geonode.ModuleName,
		Short:                      "tx commands for the geonode module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxCreateCmd(),
		TxUpdateCuratorCmd(),
		TxUpdateMetadataCmd(),
	)

	return cmd
}

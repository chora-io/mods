package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/chora-io/mods/content"
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
		TxCreateContentCmd(),
		TxRemoveContentCmd(),
		TxUpdateContentCuratorCmd(),
		TxUpdateContentMetadataCmd(),
	)

	return cmd
}

package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/choraio/mods/example"
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
		TxCreateContentCmd(),
		TxUpdateContentCmd(),
		TxDeleteContentCmd(),
	)

	return cmd
}

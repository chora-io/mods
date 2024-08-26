package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/chora-io/mods/admin"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        admin.ModuleName,
		Short:                      "tx commands for the admin module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxUpdateAdminCmd(),
	)

	return cmd
}

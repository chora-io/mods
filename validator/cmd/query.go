package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/choraio/mods/validator"
)

// QueryCmd creates and returns the query command.
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        validator.ModuleName,
		Short:                      "query commands for the validator module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryMaxMissedBlocksCmd(),
		QueryValidatorCmd(),
		QueryValidatorsCmd(),
	)

	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/chora-io/mods/validator"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        validator.ModuleName,
		Short:                      "tx commands for the validator module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxCreateValidatorCmd(),
		TxRemoveValidatorCmd(),
		TxUpdateValidatorCmd(),
	)

	return cmd
}

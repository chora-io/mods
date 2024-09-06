package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/chora-io/mods/subject"
)

// TxCmd creates and returns the tx command.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        subject.ModuleName,
		Short:                      "tx commands for the subject module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		TxCreateSubjectCmd(),
		TxRemoveSubjectCmd(),
		TxUpdateSubjectMetadataCmd(),
		TxUpdateSubjectStewardCmd(),
	)

	return cmd
}

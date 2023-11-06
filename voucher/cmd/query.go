package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/chora-io/mods/voucher"
)

// QueryCmd creates and returns the query command.
func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        voucher.ModuleName,
		Short:                      "query commands for the voucher module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryBalanceCmd(),
		QueryBalancesByAddressCmd(),
		QueryBalancesByVoucherCmd(),
		QueryVoucherCmd(),
		QueryVouchersCmd(),
		QueryVouchersByIssuerCmd(),
	)

	return cmd
}

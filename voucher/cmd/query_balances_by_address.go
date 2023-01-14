package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// QueryBalancesByAddressCmd creates and returns the query balances-by-address command.
func QueryBalancesByAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances-by-address [address]",
		Short: "query total balances by address",
		Long:  "query total balances by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			pgn, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := v1.QueryBalancesByAddressRequest{
				Address:    args[0],
				Pagination: pgn,
			}

			res, err := c.BalancesByAddress(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "balances-by-address")

	return cmd
}

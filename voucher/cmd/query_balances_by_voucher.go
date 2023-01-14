package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// QueryBalancesByVoucherCmd creates and returns the query balances-by-voucher command.
func QueryBalancesByVoucherCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances-by-voucher [voucher]",
		Short: "query total balances by voucher",
		Long:  "query total balances by voucher",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			pgn, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := v1.QueryBalancesByVoucherRequest{
				Id:         id,
				Pagination: pgn,
			}

			res, err := c.BalancesByVoucher(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "balances-by-voucher")

	return cmd
}

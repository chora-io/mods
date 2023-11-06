package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/chora-io/mods/voucher/types/v1"
)

// QueryBalanceCmd creates and returns the query balance command.
func QueryBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance [id] [address]",
		Short: "query balance by id and address",
		Long:  "query balance by id and address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			req := v1.QueryBalanceRequest{
				Id:      id,
				Address: args[1],
			}

			res, err := c.Balance(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

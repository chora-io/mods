package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/chora-io/mods/voucher/types/v1"
)

// QueryVoucherCmd creates and returns the query voucher command.
func QueryVoucherCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "voucher [id]",
		Short: "query voucher by id",
		Long:  "query voucher by id",
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

			req := v1.QueryVoucherRequest{
				Id: id,
			}

			res, err := c.Voucher(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

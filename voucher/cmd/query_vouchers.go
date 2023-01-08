package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// QueryVouchersCmd creates and returns the query vouchers command.
func QueryVouchersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vouchers",
		Short: "query all vouchers",
		Long:  "query all vouchers",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryVouchersRequest{}

			res, err := c.Vouchers(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

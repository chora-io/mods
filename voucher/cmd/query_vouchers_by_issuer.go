package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// QueryVouchersByIssuerCmd creates and returns the query vouchers-by-issuer command.
func QueryVouchersByIssuerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vouchers-by-issuer [issuer]",
		Short: "query vouchers by issuer",
		Long:  "query vouchers by issuer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryVouchersByIssuerRequest{
				Issuer: args[0],
			}

			res, err := c.VouchersByIssuer(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

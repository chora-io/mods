package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// QueryPolicyCmd creates and returns the query validator command.
func QueryPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "max-missed-blocks",
		Short: "query the maximum number of missed blocks",
		Long:  "query the maximum number of missed blocks",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			res, err := c.Policy(cmd.Context(), &v1.QueryPolicyRequest{})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

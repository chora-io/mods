package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/content/types/v1"
)

// QueryContentByCuratorCmd creates and returns the query content-by-curator command.
func QueryContentByCuratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content-by-curator [curator]",
		Short: "query content by the curator of the content",
		Long:  "query content by the curator of the content",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryContentByCuratorRequest{
				Curator: args[0],
			}

			res, err := c.ContentByCurator(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

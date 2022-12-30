package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/content/types/v1"
)

// QueryContentsByCuratorCmd creates and returns the query contents-by-curator command.
func QueryContentsByCuratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contents-by-curator [curator]",
		Short: "query contents by curator",
		Long:  "query contents by curator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryContentsByCuratorRequest{
				Curator: args[0],
			}

			res, err := c.ContentsByCurator(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

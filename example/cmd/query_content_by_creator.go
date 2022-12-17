package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/choraio/mods/example/types/v1"
)

// QueryContentByCreatorCmd creates and returns the query content-by-creator command.
func QueryContentByCreatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content-by-creator [creator]",
		Short: "query content by the creator of the content",
		Long:  "query content by the creator of the content",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryContentByCreatorRequest{
				Creator: args[0],
			}

			res, err := c.ContentByCreator(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

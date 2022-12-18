package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/example/types/v1"
)

// QueryContentCmd creates and returns the query content command.
func QueryContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content [id]",
		Short: "query content by the unique identifier of the content",
		Long:  "query content by the unique identifier of the content",
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

			req := v1.QueryContentRequest{
				Id: id,
			}

			res, err := c.Content(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

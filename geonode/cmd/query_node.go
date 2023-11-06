package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/chora-io/mods/geonode/types/v1"
)

// QueryNodeCmd creates and returns the query node command.
func QueryNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node [id]",
		Short: "query node by id",
		Long:  "query node by id",
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

			req := v1.QueryNodeRequest{
				Id: id,
			}

			res, err := c.Node(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

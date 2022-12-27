package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/geonode/types/v1"
)

// QueryNodeByCuratorCmd creates and returns the query node-by-curator command.
func QueryNodeByCuratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-by-curator [curator]",
		Short: "query node by the curator of the node",
		Long:  "query node by the curator of the node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryNodeByCuratorRequest{
				Curator: args[0],
			}

			res, err := c.NodeByCurator(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

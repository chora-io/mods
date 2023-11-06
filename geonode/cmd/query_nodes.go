package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/chora-io/mods/geonode/types/v1"
)

// QueryNodesCmd creates and returns the query nodes command.
func QueryNodesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodes",
		Short: "query all nodes",
		Long:  "query all nodes",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			pgn, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := v1.QueryNodesRequest{
				Pagination: pgn,
			}

			res, err := c.Nodes(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nodes")

	return cmd
}

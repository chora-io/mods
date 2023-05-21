package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// QueryValidatorsCmd creates and returns the query validators command.
func QueryValidatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators",
		Short: "query all validators",
		Long:  "query all validators",
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

			req := v1.QueryValidatorsRequest{
				Pagination: pgn,
			}

			res, err := c.Validators(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "validators")

	return cmd
}

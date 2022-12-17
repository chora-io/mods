package client

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/choraio/mods/example"
	types "github.com/choraio/mods/example/types/v1"
)

// GetQueryCmd creates and returns the query command.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        example.ModuleName,
		Short:                      "query commands for the example module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		QueryContentCmd(),
		QueryContentByCreatorCmd(),
	)

	return cmd
}

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

			req := types.QueryContentRequest{
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

			req := types.QueryContentByCreatorRequest{
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

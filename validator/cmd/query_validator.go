package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// QueryValidatorCmd creates and returns the query validator command.
func QueryValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator [address]",
		Short: "query validator by address",
		Long:  "query validator by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := getQueryClient(cmd)
			if err != nil {
				return err
			}

			req := v1.QueryValidatorRequest{
				Address: args[0],
			}

			res, err := c.Validator(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

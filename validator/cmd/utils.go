package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	v1 "github.com/choraio/mods/validator/types/v1"
)

func getQueryClient(cmd *cobra.Command) (v1.QueryClient, client.Context, error) {
	ctx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, client.Context{}, err
	}
	return v1.NewQueryClient(ctx), ctx, err
}

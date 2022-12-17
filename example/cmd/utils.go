package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	types "github.com/choraio/mods/example/types/v1"
)

func getQueryClient(cmd *cobra.Command) (types.QueryClient, client.Context, error) {
	ctx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, client.Context{}, err
	}
	return types.NewQueryClient(ctx), ctx, err
}

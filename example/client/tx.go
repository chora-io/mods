package client

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/choraio/mods/example"
	types "github.com/choraio/mods/example/types/v1"
)

// GetTxCmd creates and returns the tx command.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        example.ModuleName,
		Short:                      "tx commands for the example module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CreateContentCmd(),
		UpdateContentCmd(),
		DeleteContentCmd(),
	)

	return cmd
}

func CreateContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-content [hash]",
		Short: "submit a transaction to create content",
		Long:  "submit a transaction to create content",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgCreateContent{
				Creator: clientCtx.GetFromAddress().String(),
				Hash:    args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func UpdateContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-content [id] [new-hash]",
		Short: "submit a transaction to update content",
		Long:  "submit a transaction to update content",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateContent{
				Id:      id,
				Creator: clientCtx.GetFromAddress().String(),
				NewHash: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func DeleteContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-content [id]",
		Short: "submit a transaction to delete content",
		Long:  "submit a transaction to delete content",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			msg := types.MsgDeleteContent{
				Id:      id,
				Creator: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

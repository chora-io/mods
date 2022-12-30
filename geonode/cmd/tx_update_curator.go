package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/choraio/mods/geonode/types/v1"
)

// TxUpdateCuratorCmd creates and returns the tx update-curator command.
func TxUpdateCuratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-curator [id] [new-curator]",
		Short: "submit a transaction to update node curator",
		Long:  "submit a transaction to update node curator",
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

			msg := v1.MsgUpdateCurator{
				Id:         id,
				Curator:    clientCtx.GetFromAddress().String(),
				NewCurator: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/chora-io/mods/voucher/types/v1"
)

// TxUpdateMetadataCmd creates and returns the tx update-metadata command.
func TxUpdateMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-metadata [id] [new-metadata]",
		Short: "submit transaction to update voucher metadata",
		Long:  "submit transaction to update voucher metadata",
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

			msg := v1.MsgUpdateMetadata{
				Id:          id,
				Issuer:      clientCtx.GetFromAddress().String(),
				NewMetadata: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

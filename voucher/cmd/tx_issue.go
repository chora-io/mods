package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// TxIssueCmd creates and returns the tx issue command.
func TxIssueCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue [id] [recipient] [amount] [expiration] [metadata]",
		Short: "submit a transaction to issue vouchers",
		Long:  "submit a transaction to issue vouchers",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			expiration, err := time.Parse(time.RFC3339, args[3])
			if err != nil {
				return err
			}

			msg := v1.MsgIssue{
				Id:         id,
				Issuer:     clientCtx.GetFromAddress().String(),
				Recipient:  args[1],
				Amount:     args[2],
				Expiration: &expiration,
				Metadata:   args[4],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

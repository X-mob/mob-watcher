package cmd

import (
	"fmt"

	"github.com/X-mob/mob-watcher/db"
	"github.com/spf13/cobra"
)

var getMobCmd = &cobra.Command{
	Use:   "get-mob",
	Short: "get mobs detail",
	Long:  `get mob detail in db`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		mob := db.GetMob(address)

		fmt.Printf("name: %s\n", mob.Name)
		fmt.Printf("address: %s\n", mob.Address.Hex())
		fmt.Printf("creator: %s\n", mob.Creator.Hex())
		fmt.Printf("token: %s\n", mob.Token.Hex())
		fmt.Printf("tokenId: %s\n", mob.TokenId)
		fmt.Printf("raisedTotal: %s\n", mob.RaisedTotal)
		fmt.Printf("amountTotal: %s\n", mob.AmountTotal)
		fmt.Printf("takeProfitPrice: %s\n", mob.TakeProfitPrice)
		fmt.Printf("stopLossPrice: %s\n", mob.StopLossPrice)
		fmt.Printf("deadline: %s\n", mob.Deadline)
		fmt.Printf("raisedAmountDeadLine: %s\n", mob.RaisedAmountDeadline)
		fmt.Printf("Balance: %s\n", mob.Balance)
		fmt.Printf("CreatedTime: %d\n", mob.CreatedTime)
		fmt.Printf("CreatedTime: %d\n", mob.Status)
	},
}

func init() {
	rootCmd.AddCommand(getMobCmd)

	getMobCmd.Flags().StringP("address", "a", "", "mob address")
}

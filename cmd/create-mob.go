package cmd

import (
	"time"

	"github.com/X-mob/mob-watcher/lib"
	"github.com/spf13/cobra"
)

var createMobCmd = &cobra.Command{
	Use:   "create-mob",
	Short: "create an new mob",
	Long:  `create an new mob via contract`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		token, _ := cmd.Flags().GetString("token")
		tokenId, _ := cmd.Flags().GetInt64("tokenId")
		raisedTotal, _ := cmd.Flags().GetString("raisedTotal")
		takeProfitPrice, _ := cmd.Flags().GetString("takeProfitPrice")
		stopLossPrice, _ := cmd.Flags().GetString("stopLossPrice")
		raisedAmountDeadline, _ := cmd.Flags().GetInt64("raisedAmountDeadline")
		deadline, _ := cmd.Flags().GetInt64("deadline")

		lib.CreateMob(name, token, tokenId, raisedTotal, takeProfitPrice, stopLossPrice, raisedAmountDeadline, deadline)
	},
}

func init() {
	rootCmd.AddCommand(createMobCmd)

	createMobCmd.Flags().StringP("name", "n", "Default Mob", "mob name")
	createMobCmd.Flags().StringP("token", "t", "", "mob nft token address")
	createMobCmd.Flags().Int64P("tokenId", "i", 0, "mob nft token id")
	createMobCmd.Flags().StringP("raisedTotal", "r", "0", "mob total raise fund target")
	createMobCmd.Flags().StringP("takeProfitPrice", "p", "0", "mob take profit price")
	createMobCmd.Flags().StringP("stopLossPrice", "s", "0", "mob stop loss price")
	createMobCmd.Flags().Int64P("raisedAmountDeadline", "a", time.Now().AddDate(0, 3, 0).Unix(), "mob deadline for fund raising")
	createMobCmd.Flags().Int64P("deadline", "d", time.Now().AddDate(0, 6, 0).Unix(), "mob deadline for all")
}

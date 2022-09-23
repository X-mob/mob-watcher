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
		raiseTarget, _ := cmd.Flags().GetString("raiseTarget")
		takeProfitPrice, _ := cmd.Flags().GetString("takeProfitPrice")
		stopLossPrice, _ := cmd.Flags().GetString("stopLossPrice")
		raiseDeadline, _ := cmd.Flags().GetUint64("raiseDeadline")
		deadline, _ := cmd.Flags().GetUint64("deadline")
		targetMode, _ := cmd.Flags().GetUint8("targetMode")

		if targetMode > 1 {
			panic("0 for restrictMode, 1 for fullOpen mode")
		}

		lib.CreateMob(name, token, tokenId, raiseTarget, takeProfitPrice, stopLossPrice, raiseDeadline, deadline, targetMode)
	},
}

func init() {
	rootCmd.AddCommand(createMobCmd)

	createMobCmd.Flags().StringP("name", "n", "Default Mob", "mob name")
	createMobCmd.Flags().StringP("token", "t", "", "mob nft token address")
	createMobCmd.Flags().Int64P("tokenId", "i", 0, "mob nft token id")
	createMobCmd.Flags().StringP("raiseTarget", "r", "0", "mob total raise fund target")
	createMobCmd.Flags().StringP("takeProfitPrice", "p", "0", "mob take profit price")
	createMobCmd.Flags().StringP("stopLossPrice", "s", "0", "mob stop loss price")
	createMobCmd.Flags().Uint64P("raiseDeadline", "a", uint64(time.Now().AddDate(0, 3, 0).Unix()), "mob deadline for fund raising")
	createMobCmd.Flags().Uint64P("deadline", "d", uint64(time.Now().AddDate(0, 6, 0).Unix()), "mob deadline for all")
	createMobCmd.Flags().Uint8P("targetMode", "m", 0, "NFT tokenId TargetMode")
}

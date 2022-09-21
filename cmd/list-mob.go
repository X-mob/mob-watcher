package cmd

import (
	"fmt"

	"github.com/X-mob/mob-watcher/db"
	"github.com/spf13/cobra"
)

var listMobCmd = &cobra.Command{
	Use:   "list-mob",
	Short: "list all mobs address",
	Long:  `list all mobs address in db`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetInt64("status")
		if status == -1 { // -1 means no filtering
			// todo: apply limit arg
			addrs := db.GetAllMobAddress()
			fmt.Println(addrs)
		} else {
			addrs := db.GetMobAddressByStatus(db.MobStatus(status))
			fmt.Println(addrs)
		}
	},
}

func init() {
	rootCmd.AddCommand(listMobCmd)

	listMobCmd.Flags().Int64P("limit", "l", -1, "limit number")
	listMobCmd.Flags().Int64P("status", "s", -1, "Mob Status to filter")
}

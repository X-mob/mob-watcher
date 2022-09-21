package cmd

import (
	"github.com/X-mob/mob-watcher/lib"
	"github.com/spf13/cobra"
)

var joinMobCmd = &cobra.Command{
	Use:   "join-mob",
	Short: "join an new mob",
	Long:  `join an new mob via contract`,
	Run: func(cmd *cobra.Command, args []string) {
		value, _ := cmd.Flags().GetString("value")
		address, _ := cmd.Flags().GetString("address")

		lib.JoinMob(value, address)
	},
}

func init() {
	rootCmd.AddCommand(joinMobCmd)

	joinMobCmd.Flags().StringP("value", "v", "0", "join deposit value")
	joinMobCmd.Flags().StringP("address", "a", "", "mob address")
}

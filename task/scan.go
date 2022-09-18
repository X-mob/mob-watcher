package task

import (
	"fmt"

	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
)

func ScanNewMob() {
	events := lib.GetMobsCreate(nil)
	for _, event := range events {
		fmt.Println(event)
		fmt.Println(db.GetKV("answer"))
	}

}

func ScanNewJoin() {

}

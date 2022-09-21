package task

import (
	"fmt"
	"time"
)

func Start() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fmt.Println("scanning..")
			ScanNewMob()
			ScanRaisingMob()
			ScanRaiseFinishedMob()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

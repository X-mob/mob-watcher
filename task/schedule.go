package task

import (
	"time"
)

func Start() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			Scan()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

package task

import "time"

func start() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

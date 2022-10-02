package task

import (
	"time"
)

var GlobalTxManager = NewTxManager()

func Start() {
	go func() {
		TickerJob()
	}()

	// single isolate job
	SingeJob(ScanRaisingMob)
	SingeJob(ScanRaiseSuccessMob)
	SingeJob(ScanBuyFailedMob)
	SingeJob(ScanNftBoughtMob)
	SingeJob(ScanNftSoldMob)
	SingeJob(ScanCanClaimMob)

	// other process
	SingeJob(ScanNewJoin)
	SingeJob(ScanNftBalanceAttacking)
}

func TickerJob() {
	ticker := time.NewTicker(10 * time.Second)
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

func SingeJob(f func()) {
	go func() {
		RecurLoop(f)
	}()
}

func RecurLoop(f func()) {
	f()
	time.Sleep(5 * time.Second) // wait 5 secs
	RecurLoop(f)
}

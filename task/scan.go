package task

import (
	"fmt"
	"time"

	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/utils"
)

func ScanNewMob() {
	var cursor uint64

	// get cursor
	cursorByte := db.GetKVWithBytes(db.SCAN_MOB_CURSOR_KEY)
	if len(cursorByte) == 0 {
		cursor = 0
	} else {
		cursor = utils.StringToBigInt(string(cursorByte[:])).Uint64()
	}

	events, nextCursor := lib.GetMobsCreate(cursor, nil)

	// save new cursor
	db.SetKV(db.SCAN_MOB_CURSOR_KEY, fmt.Sprint(*nextCursor))

	for _, event := range events {
		mob := db.MobCreateToMob(event)
		isExit := db.IsMobExits(mob.Address.Hex())
		if isExit == false {
			db.SetMob(mob)
			fmt.Println("save mob total: ", len(events))

			// init patch
			lib.PatchMobWithWethSeaport(mob.Address.Hex())
		}
	}
}

func ScanRaisingMob() {
	raisingMob := db.GetMobsWithStatus(db.Raising)
	for _, m := range raisingMob {
		mob, err := lib.NewXmobExchangeCore(m.Address, lib.EthClient)
		panicErr(err)

		amount, _ := mob.AmountTotal(nil)
		target, _ := mob.RaisedTotal(nil)
		fmt.Println(amount, target, amount.Uint64() == target.Uint64())
		if amount.Uint64() == target.Uint64() {
			// update mob status
			fmt.Println("update mob status", m.Address.Hex(), db.RaiseFailed)
			db.UpdateMobStatus(m.Address.Hex(), db.RaiseFailed)
			db.UpdateMobAmountTotal(m.Address.Hex(), *amount)
			return
		}

		if m.RaisedAmountDeadline.Int64() <= time.Now().Unix() {
			// update mob status
			fmt.Println("update mob status", m.Address.Hex(), db.RaiseFailed)
			db.UpdateMobStatus(m.Address.Hex(), db.RaiseFailed)
			return
		}

		if amount.Uint64() != m.AmountTotal.Uint64() {
			fmt.Println("update mob amount total", m.Address.Hex(), amount)
			db.UpdateMobAmountTotal(m.Address.Hex(), *amount)
			return
		}

		fmt.Println("no update for raising mob..", m.Address.Hex())
	}
}

func ScanRaiseFinishedMob() {
	raiseFinishedMob := db.GetMobsWithStatus(db.RaiseFinished)
	for _, m := range raiseFinishedMob {
		buyNowEvents := lib.GetBuyNow(m.Address.Hex())
		if len(buyNowEvents) != 0 {
			// update status
			db.UpdateMobStatus(m.Address.Hex(), db.NftBought)
			return
		}

		// trigger buy
		// mob := lib.GetMobByAddress(m.Address.Hex())
		// mob.BuyNow(lib.BasicTransactionOpts, )

		fmt.Println("no update for raiseFinished mob..", m.Address.Hex())
	}
}

func ScanNftBoughtMob() {

}

func ScanNftSoldMob() {

}

func ScanCanClaimMob() {

}

func ScanNewJoin() {

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

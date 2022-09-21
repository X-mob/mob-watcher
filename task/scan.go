package task

import (
	"fmt"
	"time"

	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/utils"
)

func Scan() {
	fmt.Println("scanning..")

	ScanNewMob()
	ScanRaisingMob()
	ScanRaiseFinishedMob()
}

func ScanNewMob() {
	var cursor uint64

	// get cursor
	cursorByte := db.GetKVWithBytes(db.SCAN_MOB_CURSOR_KEY)
	if len(cursorByte) == 0 {
		cursor = 0
	} else {
		cursor = utils.StringToBigInt(string(cursorByte[:])).Uint64()
	}

	mobCreateEvents, nextCursor := lib.GetMobsCreateEvents(cursor, nil)

	// save new cursor
	db.SetKV(db.SCAN_MOB_CURSOR_KEY, fmt.Sprint(*nextCursor))

	for _, event := range mobCreateEvents {
		mob := db.MobCreateToMob(event)
		isExit := db.IsMobExits(mob.Address.Hex())
		if isExit == false {
			db.AddMob(mob)
			fmt.Println("save mob total: ", len(mobCreateEvents))

			// init patch
			// todo: remove this in mainnet
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
		if amount.Uint64() == target.Uint64() {
			// raise finish
			const status = db.RaiseFinished
			fmt.Printf("update mob status: %s, address: %s\n", status.String(), m.Address.Hex())
			db.UpdateMobStatus(m.Address.Hex(), status)
			db.UpdateMobAmountTotal(m.Address.Hex(), *amount)
			db.UpdateMobBalance(m.Address.Hex(), *amount)
			return
		}

		if m.RaisedAmountDeadline.Int64() <= time.Now().Unix() {
			// over raise deadline
			const status = db.RaiseFailed
			fmt.Printf("update mob status: %s, address: %s\n", status.String(), m.Address.Hex())
			db.UpdateMobStatus(m.Address.Hex(), status)
			return
		}

		if amount.Uint64() != m.AmountTotal.Uint64() {
			// new fund adding
			fmt.Printf("update mob amount total: %s, address: %s\n", amount, m.Address.Hex())
			db.UpdateMobAmountTotal(m.Address.Hex(), *amount)
			return
		}
	}
}

func ScanRaiseFinishedMob() {
	raiseFinishedMob := db.GetMobsWithStatus(db.RaiseFinished)
	for _, m := range raiseFinishedMob {
		buyNowEvents := lib.GetBuyNowEvents(m.Address.Hex())
		if len(buyNowEvents) != 0 {
			// already bought
			db.UpdateMobStatus(m.Address.Hex(), db.NftBought)
			return
		}

		// try buy
		BuyNow(m.Address.Hex())
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

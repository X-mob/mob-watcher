package task

import (
	"fmt"
	"time"

	"github.com/X-mob/mob-watcher/casting"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/utils"
)

func Scan() {
	fmt.Println("scanning..")

	ScanNewMob()

	go func() {
		ScanRaisingMob()
		ScanRaiseSuccessMob()
		ScanBuyFailedMob()
		ScanNftBoughtMob()
		ScanNftSoldMob()
		ScanCanClaimMob()

		// other process
		ScanNewJoin()
		ScanNftBalanceAttacking()
	}()
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
			// lib.PatchMobWithSeaport(mob.Address.Hex())
		}
	}
}

func ScanRaisingMob() {
	raisingMob := db.GetMobsWithStatus(db.Raising)
	for _, m := range raisingMob {
		mob, err := lib.NewXmobExchangeCore(m.Address, lib.EthClient)
		panicErr(err)

		metadata, _ := mob.Metadata(nil)
		amount := metadata.RaisedAmount
		target := metadata.RaiseTarget
		if amount.Uint64() == target.Uint64() {
			// raise finish
			const status = db.RaiseSuccess
			fmt.Printf("update mob status: %s, address: %s\n", status.String(), m.Address.Hex())
			db.UpdateMobStatus(m.Address.Hex(), status)
			db.UpdateMobRaisedAmount(m.Address.Hex(), *amount)
			db.UpdateMobBalance(m.Address.Hex(), *amount)
			return
		}

		if m.RaiseDeadline.Int64() < time.Now().Unix() {
			// over raise deadline
			const status = db.RaiseFailed
			fmt.Printf("update mob status: %s, address: %s\n", status.String(), m.Address.Hex())
			db.UpdateMobStatus(m.Address.Hex(), status)
			return
		}

		if amount.Uint64() != m.RaisedAmount.Uint64() {
			// update with new fund adding
			fmt.Printf("update mob amount total: %s, address: %s\n", amount, m.Address.Hex())
			db.UpdateMobRaisedAmount(m.Address.Hex(), *amount)
			return
		}
	}
}

func ScanRaiseSuccessMob() {
	raiseSuccessMob := db.GetMobsWithStatus(db.RaiseSuccess)
	for _, m := range raiseSuccessMob {
		buyNowEvents := lib.GetBuyNowEvents(m.Address.Hex())
		if len(buyNowEvents) != 0 {
			// already bought
			db.UpdateMobStatus(m.Address.Hex(), db.NftBought)
			return
		}

		if m.Deadline.Int64() > time.Now().Unix() {
			// buy failed
			db.UpdateMobStatus(m.Address.Hex(), db.NftBuyFailed)
			return
		}

		// try to buy
		fmt.Printf("try buying, mob: %s", m.Address.Hex())
		BuyNow(m.Address.Hex())
	}
}

func ScanBuyFailedMob() {
	nftBuyFailedMob := db.GetMobsWithStatus(db.NftBuyFailed)
	for _, m := range nftBuyFailedMob {
		fmt.Printf("try settling after buy failed, mob: %s", m.Address.Hex())
		lib.SettleAfterBuyFailed(m.Address.Hex())
	}
}

func ScanNftBoughtMob() {
	nftBoughtMob := db.GetMobsWithStatus(db.NftBought)
	for _, m := range nftBoughtMob {
		if lib.IsNFTUnOwned(m.Address.Hex()) == true {
			fmt.Printf("mob %s has no nft, already sold?", m.Address.Hex())
			// already sold?
			db.UpdateMobStatus(m.Address.Hex(), db.NftSold)
			return
		}

		if m.Deadline.Int64() > time.Now().Unix() {
			// over deadline, try settle to refund for users
			fmt.Printf("try settling after deadline, mob: %s", m.Address.Hex())
			lib.SettleAfterDeadline(m.Address.Hex())
		}

		// try to pre-sell at take profit price
		// todo: check with floor price to determine if register more sell order
		// todo: add some strategy
		Sell(m.Address.Hex(), TakeProfit)

		// todo: check real floor price to support stop loss selling
		// slug := "" // db save metadata of nft
		// res := opensea.GetCollectionStats(slug)
		// if res.Stats.FloorPrice < m.StopLossPrice {
		// 	Sell(m.Address.Hex(), StopLoss)
		// }
	}
}

func ScanNftSoldMob() {
	nftSoldMob := db.GetMobsWithStatus(db.NftSold)
	for _, m := range nftSoldMob {
		settlementEvents := lib.GetSettlementEvents(m.Address.Hex())
		if len(settlementEvents) != 0 {
			// already settle
			db.UpdateMobStatus(m.Address.Hex(), db.CanClaim)
			return
		}

		settlementAfterDeadlineEvents := lib.GetSettlementAfterDeadlineEvents(m.Address.Hex())
		if len(settlementAfterDeadlineEvents) != 0 {
			// already settle
			db.UpdateMobStatus(m.Address.Hex(), db.CanClaim)
			return
		}

		// try to settle
		fmt.Printf("try settling, mob: %s", m.Address.Hex())
		lib.Settle(m.Address.Hex())
	}
}

func ScanCanClaimMob() {
	canClaimMob := db.GetMobsWithStatus(db.CanClaim)
	for _, m := range canClaimMob {
		xmobStatus := lib.GetMobStatus(m.Address.Hex())
		status := casting.XmobToDBMobStatus(xmobStatus)
		if status == db.AllClaimed {
			// all claimed
			db.UpdateMobStatus(m.Address.Hex(), status)
		}
	}
}

func ScanNewJoin() {
	// todo
}

func ScanNftBalanceAttacking() {
	// todo
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

package casting

import (
	"github.com/X-mob/mob-watcher/db"
)

// MobStatus from Xmob contract
// enum MobStatus {
// 	RAISING = 0,
// 	RAISE_SUCCESS,
// 	NFT_BOUGHT,
// 	CAN_CLAIM, // NFT_SOLD is not including, we can not check nft is sold on chain
// 	ALL_CLAIMED
// }

func XmobToDBMobStatus(status uint8) db.MobStatus {
	switch status {
	case 0:
		return db.Raising
	case 1:
		return db.RaiseSuccess
	case 2:
		return db.NftBought
	case 3:
		return db.CanClaim
	case 4:
		return db.AllClaimed

	default:
		panic("unknown xmob status")
	}
}

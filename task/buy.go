package task

import (
	"fmt"

	"github.com/X-mob/mob-watcher/casting"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/opensea"
)

func GetOrderFromOpenSea(token string, tokenId string) *lib.Order {
	var opt opensea.RetrieveOffersOption
	// target the nft
	opt.AssetContractAddress = token
	opt.TokenIds = []string{tokenId}
	// get cheapest one
	opt.Limit = 1
	opt.OrderBy = "eth_price"
	opt.OrderDirection = "asc"

	res := opensea.RetrieveOffers(opt)

	if len(res.Orders) == 0 {
		return nil
	}

	order := casting.OpenSeaToSeaportOrder(res.Orders[0].ProtocolData)
	return &order
}

// todo: update token id after bought under full-open target mode..
func BuyNow(mobAddress string) {
	mobInDB := db.GetMob(mobAddress)
	tokenAddress := mobInDB.Token.Hex()
	tokenId := mobInDB.TokenId.String()
	order := GetOrderFromOpenSea(tokenAddress, tokenId)
	if order == nil {
		fmt.Println("order not found from OpenSea, skip")
		return
	}

	lib.BuyOrder(mobAddress, *order)
}

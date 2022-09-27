package task

import (
	"fmt"

	"github.com/X-mob/mob-watcher/casting"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/opensea"
)

func GetListingFromOpenSea(token string, tokenId string) *lib.Order {
	var opt opensea.RetrieveListingsOption
	// target the nft
	opt.AssetContractAddress = token
	opt.TokenIds = []string{tokenId}

	// todo: only opensea mainnet api support the following params
	opt.Limit = 1
	// opt.OrderBy = "eth_price"
	// opt.OrderDirection = "asc"

	res := opensea.RetrieveListings(opt)

	if len(res.Orders) == 0 {
		return nil
	}

	order := casting.OpenSeaToSeaportOrder(res.Orders[0].ProtocolData)
	return &order
}

func GetOfferFromOpenSea(token string, tokenId string) *lib.Order {
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

func BuyNow(mobAddress string) {
	mobInDB := db.GetMob(mobAddress)
	tokenAddress := mobInDB.Token.Hex()
	tokenId := mobInDB.TokenId.String()
	order := GetListingFromOpenSea(tokenAddress, tokenId)
	if order == nil {
		fmt.Println("order not found from OpenSea, skip")
		return
	}

	fmt.Printf("%+v\n", order)

	lib.BuyOrder(mobAddress, *order)
}

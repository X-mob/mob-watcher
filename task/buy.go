package task

import (
	"encoding/hex"
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

	//fmt.Printf("%+v\n", res.Orders[0].ProtocolData)
	//AssertOrderCasting(order)

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
		fmt.Println("order not found from OpenSea, skip..")
		return
	}

	lib.BuyOrder(mobAddress, *order)
}

func AssertOrderCasting(order lib.Order) {
	conduitKey := order.Parameters.ConduitKey
	fmt.Printf("conduitKey: %s\n", hex.EncodeToString(conduitKey[:]))

	zoneHash := order.Parameters.ZoneHash
	fmt.Printf("zoneHash: %s\n", hex.EncodeToString(zoneHash[:]))

	orderType := order.Parameters.OrderType
	fmt.Printf("orderType: %d\n", orderType)

	offerer := order.Parameters.Offerer
	fmt.Printf("offerer: %s\n", offerer.Hex())

	zone := order.Parameters.Zone
	fmt.Printf("zone: %s\n", zone.Hex())

	startTime := order.Parameters.StartTime
	fmt.Printf("startTime: %s\n", startTime.String())

	endTime := order.Parameters.EndTime
	fmt.Printf("endTime: %s\n", endTime.String())

	salt := order.Parameters.Salt
	fmt.Printf("salt: %s\n", salt.String())

	totalOriginalConsiderationItems := order.Parameters.TotalOriginalConsiderationItems
	fmt.Printf("TotalOriginalConsiderationItems: %s\n", totalOriginalConsiderationItems.String())

	fmt.Printf("offer ==> length: %d\n", len(order.Parameters.Offer))
	AssertOffer(order.Parameters.Offer[0])

	fmt.Printf("consideration ==>  length: %d\n", len(order.Parameters.Consideration))
	AssertConsideration(order.Parameters.Consideration[0])

	sig := order.Signature
	fmt.Printf("signature: %s\n", hex.EncodeToString(sig))
}

func AssertOffer(offer lib.OfferItem) {
	itemType := offer.ItemType
	fmt.Printf("itemType: %d\n", itemType)

	token := offer.Token
	fmt.Printf("token: %s\n", token.Hex())

	id := offer.IdentifierOrCriteria
	fmt.Printf("IdentifierOrCriteria: %s\n", id.String())

	start := offer.StartAmount
	fmt.Printf("StartAmount: %s\n", start.String())

	end := offer.EndAmount
	fmt.Printf("EndAmount: %s\n", end.String())
}

func AssertConsideration(consider lib.ConsiderationItem) {
	itemType := consider.ItemType
	fmt.Printf("itemType: %d\n", itemType)

	token := consider.Token
	fmt.Printf("token: %s\n", token.Hex())

	id := consider.IdentifierOrCriteria
	fmt.Printf("IdentifierOrCriteria: %s\n", id.String())

	start := consider.StartAmount
	fmt.Printf("StartAmount: %s\n", start.String())

	end := consider.EndAmount
	fmt.Printf("EndAmount: %s\n", end.String())

	recipient := consider.Recipient
	fmt.Printf("Recipient: %s\n", recipient.Hex())
}

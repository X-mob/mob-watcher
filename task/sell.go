package task

import (
	"fmt"
	"math/big"
	"time"

	"github.com/X-mob/mob-watcher/casting"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/opensea"
	"github.com/X-mob/mob-watcher/utils"
	"github.com/ethereum/go-ethereum/common"
)

type SellType int64

const (
	TakeProfit SellType = iota
	StopLoss            // the deadline is reached but raise not enough
)

func (s SellType) String() string {
	switch s {
	case TakeProfit:
		return "TakeProfit"
	case StopLoss:
		return "RaiseFailed"
	}
	return "unknown"
}

func Sell(mobAddress string, sellType SellType) {
	_order := BuildSellOrder(mobAddress, sellType)
	fmt.Printf("%+v\n", _order)
	order := casting.OpenSeaToSeaportOrder(_order)
	lib.RegisterSellOrder(mobAddress, []lib.Order{order})
	PostOrderToOpenSea(_order)
}

// todo: introduce diff params to keep selling when not sold out
func BuildSellOrder(mobAddress string, sellType SellType) opensea.ProtocolData {
	var order opensea.ProtocolData

	mobInDB := db.GetMob(mobAddress)

	startTime := big.NewInt(time.Now().Unix())
	endTime := big.NewInt(time.Now().AddDate(0, 6, 0).Unix())
	offerer := common.HexToAddress(mobAddress)
	counter := GetCounter(offerer)
	ethPrice := mobInDB.TakeProfitPrice

	offer := opensea.OfferItem{
		ItemType:             2, // ERC721
		Token:                mobInDB.Token.Hex(),
		IdentifierOrCriteria: mobInDB.TokenId.String(),
		StartAmount:          "1",
		EndAmount:            "1",
	}
	consideration := opensea.ConsiderationItem{
		ItemType:             0, // NATIVE
		Token:                utils.ZeroAddress().Hex(),
		IdentifierOrCriteria: "0",
		StartAmount:          ethPrice.String(),
		EndAmount:            ethPrice.String(),
		Recipient:            mobAddress,
	}

	consideration2 := opensea.ConsiderationItem{
		ItemType:             0, // NATIVE
		Token:                utils.ZeroAddress().Hex(),
		IdentifierOrCriteria: "0",
		StartAmount:          ethPrice.String(),
		EndAmount:            ethPrice.String(),
		Recipient:            "0x0000a26b00c1F0DF003000390027140000fAa719",
	}

	parameters := opensea.OrderParameters{
		Offerer:                         offerer,
		Zone:                            utils.ZeroAddress(),
		Offer:                           []opensea.OfferItem{offer},
		Consideration:                   []opensea.ConsiderationItem{consideration, consideration2},
		OrderType:                       0, // FULL_OPEN
		StartTime:                       startTime.String(),
		EndTime:                         endTime.String(),
		ZoneHash:                        utils.Zero32BytesHexString(),
		Salt:                            utils.GenRandomSalt(32).String(),
		ConduitKey:                      "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
		TotalOriginalConsiderationItems: 1,
		Counter:                         int(counter.Int64()),
	}
	order = opensea.ProtocolData{
		Parameters: parameters,
		Signature:  utils.MAGIC_SIG_STR,
	}
	return order
}

func PostOrderToOpenSea(order opensea.ProtocolData) {
	opensea.CreateListing(order.Parameters, string(order.Signature))
}

func GetCounter(address common.Address) *big.Int {
	counter, err := lib.SeaportInstance.GetCounter(nil, address)
	if err != nil {
		panic(err)
	}
	return counter
}

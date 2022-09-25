package lib

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/X-mob/mob-watcher/config"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetMobsCreateEvents(start uint64, creator []common.Address) ([]XmobManageMobCreate, *uint64) {
	nextCursor := GetLatestBlockNum()
	opts := bind.FilterOpts{Start: start, End: nextCursor, Context: context.Background()}
	iterator, err := XmobManageInstance.FilterMobCreate(&opts, creator, nil, nil)
	if err != nil {
		panic(err)
	}
	var events []XmobManageMobCreate
	for iterator.Next() {
		if iterator.Event != nil {
			fmt.Printf("creator: %v\n", iterator.Event.Creator)
			fmt.Printf("token: %v\n", iterator.Event.Token)
			fmt.Printf("token id: %v\n", iterator.Event.TokenId)
		}
		events = append(events, *iterator.Event)
	}
	return events, nextCursor
}

func GetBuyNowEvents(mobAddress string) []XmobExchangeCoreBuy {
	opts := bind.FilterOpts{Start: 0, End: nil, Context: context.Background()}
	xmobExchangeCoreInstance := GetMobByAddress(mobAddress)
	iterator, err := xmobExchangeCoreInstance.FilterBuy(&opts, nil)
	if err != nil {
		panic(err)
	}
	var events []XmobExchangeCoreBuy
	for iterator.Next() {
		if iterator.Event != nil {
			fmt.Printf("seller: %v\n", iterator.Event.Seller)
		}
		events = append(events, *iterator.Event)
	}

	if len(events) > 1 {
		panic("buy event should only be fired once")
	}

	return events
}

// todo: use generic to refactor after updating go 1.8
func GetSettlementEvents(mobAddress string) []XmobExchangeCoreSettlement {
	opts := bind.FilterOpts{Start: 0, End: nil, Context: context.Background()}
	xmobExchangeCoreInstance := GetMobByAddress(mobAddress)
	iterator, err := xmobExchangeCoreInstance.FilterSettlement(&opts)
	if err != nil {
		panic(err)
	}
	var events []XmobExchangeCoreSettlement
	for iterator.Next() {
		events = append(events, *iterator.Event)
	}

	if len(events) > 1 {
		panic("settlement event should only be fired once")
	}

	return events
}

func GetSettlementAfterBuyFailedEvents(mobAddress string) []XmobExchangeCoreSettlementAfterBuyFailed {
	opts := bind.FilterOpts{Start: 0, End: nil, Context: context.Background()}
	xmobExchangeCoreInstance := GetMobByAddress(mobAddress)
	iterator, err := xmobExchangeCoreInstance.FilterSettlementAfterBuyFailed(&opts)
	if err != nil {
		panic(err)
	}
	var events []XmobExchangeCoreSettlementAfterBuyFailed
	for iterator.Next() {
		events = append(events, *iterator.Event)
	}

	if len(events) > 1 {
		panic("settlementAfterBuyFailed event should only be fired once")
	}

	return events
}

func GetSettlementAfterDeadlineEvents(mobAddress string) []XmobExchangeCoreSettlementAfterDeadline {
	opts := bind.FilterOpts{Start: 0, End: nil, Context: context.Background()}
	xmobExchangeCoreInstance := GetMobByAddress(mobAddress)
	iterator, err := xmobExchangeCoreInstance.FilterSettlementAfterDeadline(&opts)
	if err != nil {
		panic(err)
	}
	var events []XmobExchangeCoreSettlementAfterDeadline
	for iterator.Next() {
		events = append(events, *iterator.Event)
	}

	if len(events) > 1 {
		panic("settlementAfterDeadline event should only be fired once")
	}

	return events
}

func GetLogsByContract() {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(11),
		Addresses: []common.Address{
			config.GlobalConfig.XmobManagerAddress,
		},
	}

	logs, err := EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(XmobManageABI)))
	if err != nil {
		panic(err)
	}

	for _, log := range logs {

		if len(log.Data) > 0 {
			_, err := contractAbi.Unpack("MobCreate", log.Data)
			if err != nil {
				panic(err)
			}
			// fmt.Println(logInterface...)
		}

		//
		var topics [4]string
		for i := range log.Topics {
			topics[i] = log.Topics[i].Hex()
		}

	}
}

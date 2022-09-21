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

func GetMobsCreate(start uint64, creator []common.Address) ([]XmobManageMobCreate, *uint64) {
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

func GetBuyNow(mobAddress string) []XmobExchangeCoreExchanged {
	opts := bind.FilterOpts{Start: 0, End: nil, Context: context.Background()}
	xmobExchangeCoreInstance := GetMobByAddress(mobAddress)
	iterator, err := xmobExchangeCoreInstance.FilterExchanged(&opts, nil, nil)
	if err != nil {
		panic(err)
	}
	var events []XmobExchangeCoreExchanged
	for iterator.Next() {
		if iterator.Event != nil {
			fmt.Printf("buyer: %v\n", iterator.Event.Buyer)
			fmt.Printf("seller: %v\n", iterator.Event.Seller)
		}
		events = append(events, *iterator.Event)
	}

	if len(events) > 1 {
		panic("buyNow should only be fired once")
	}

	return events
}

func GetLogsByContract() {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(11),
		Addresses: []common.Address{
			config.GlobalConfig.XmobManagerContractAddress,
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
package lib

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
			events = append(events, *iterator.Event)
		}
	}
	return events, nextCursor
}

type MemberJoin struct {
	MemberAddress   common.Address
	ContractAddress common.Address
	Value           *big.Int
}

type MemberJoinData struct {
	Member common.Address
	Value  *big.Int
}

func GetMemberJoinEvents(start uint64) ([]MemberJoin, *uint64) {
	nextCursor := GetLatestBlockNum()
	eventSignature := []byte("MemberJoin(address,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(start)),
		ToBlock:   big.NewInt(int64(*nextCursor)),
		Topics: [][]common.Hash{
			{hash},
		},
	}

	logs, err := EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(XmobExchangeCoreABI)))
	if err != nil {
		panic(err)
	}

	var memberJoin []MemberJoin

	for _, log := range logs {
		data := MemberJoinData{}
		if len(log.Data) != 64 {
			continue
		}

		err := contractAbi.UnpackIntoInterface(&data, "MemberJoin", log.Data)
		if err != nil {
			panic(err)
		}

		memberJoin = append(memberJoin, MemberJoin{
			MemberAddress:   data.Member,
			ContractAddress: log.Address,
			Value:           data.Value,
		})
	}

	return memberJoin, nextCursor
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

func GetLogsByContract(query ethereum.FilterQuery) {
	logs, err := EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}
	for _, log := range logs {
		var topics [4]string
		for i := range log.Topics {
			topics[i] = log.Topics[i].Hex()
		}
	}
}

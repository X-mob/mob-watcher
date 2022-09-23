package lib

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"runtime"

	"github.com/X-mob/mob-watcher/config"
	"github.com/X-mob/mob-watcher/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetFeeRate() {
	feeRate, err := XmobManageInstance.FeeRate(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract fee rate:", feeRate)
}

func GetMobById(id *big.Int) *XmobExchangeCore {
	mobAddress, err := XmobManageInstance.MobsById(nil, id)
	if err != nil {
		panic(err)
	}
	return GetMobByAddress(mobAddress.Hex())
}

func GetMobByAddress(address string) *XmobExchangeCore {
	mob, err := NewXmobExchangeCore(common.HexToAddress(address), EthClient)
	if err != nil {
		panic(err)
	}
	return mob
}

func CreateMob(
	_name string,
	_token string,
	_tokenId int64,
	_raisedTotal string,
	_takeProfitPrice string,
	_stopLossPrice string,
	_raisedAmountDeadline int64,
	_deadline int64) {

	token := common.HexToAddress(_token)
	tokenId := big.NewInt(_tokenId)
	raisedTotal := utils.StringToBigInt(_raisedTotal)
	takeProfitPrice := big.NewInt(25)
	stopLossPrice := big.NewInt(1)
	raisedAmountDeadline := big.NewInt(_raisedAmountDeadline)
	deadline := big.NewInt(_deadline)
	mobName := _name

	txOpts := NewTxOpts(nil)
	tx, err := XmobManageInstance.CreateMob(txOpts, token, tokenId, raisedTotal, takeProfitPrice, stopLossPrice, raisedAmountDeadline, deadline, mobName)
	CheckAndWaitTx(tx, err)
}

func JoinMob(value string, mobAddress string) {
	var txOpts *bind.TransactOpts = NewTxOpts(utils.StringToBigInt(value))
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.JoinPay(txOpts, txOpts.From)
	CheckAndWaitTx(tx, err)
}

func Claim(mobAddress string) {
	mob := GetMobByAddress(mobAddress)

	// check if claim is allowed
	canClaim, err := mob.CanClaim(nil)
	if err != nil {
		panic(err)
	}
	if canClaim == false {
		panic("can not claim yet")
	}

	// claim
	txOpts := NewTxOpts(nil)
	tx, err := mob.Claim(txOpts)
	CheckAndWaitTx(tx, err)
}

func Settle(mobAddress string) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.SettlementAllocation(txOpts, false)
	CheckAndWaitTx(tx, err)
}

func BuyNow(mobAddress string, order BasicOrderParameters) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.BuyBasicOrder(txOpts, order)
	CheckAndWaitTx(tx, err)
}

func Sell(mobAddress string, orders []Order) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	tx, err := mob.RegisterSellOrder(txOpts, orders)
	CheckAndWaitTx(tx, err)
}

// only for testing, mainnet we don't have to do this
func PatchMobWithSeaport(mobAddress string) {
	txOpts := NewTxOpts(nil)
	mob := GetMobByAddress(mobAddress)
	{
		tx, err := mob.SetSeaportAddress(txOpts, config.GlobalConfig.SeaportAddress)
		CheckAndWaitTx(tx, err)
	}
}

func CheckAndWaitTx(tx *types.Transaction, txErr error) {
	// show caller info
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fmt.Printf("called from %s\n", details.Name())
	}

	// check
	if txErr != nil {
		fmt.Printf("tx failed: %s", txErr.Error())
		return
	}

	fmt.Printf("tx Hash: %s, wait to be mined..", tx.Hash())
	receipt, err := bind.WaitMined(context.Background(), EthClient, tx)
	if err != nil {
		fmt.Printf("get tx receipt failed: %s", err.Error())
		return
	}

	fmt.Printf("tx mined at block %d\n", receipt.BlockNumber.Int64())
}

func NewTxOpts(value *big.Int) *bind.TransactOpts {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)

	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	txOpts.GasPrice = gasPrice

	if value != nil {
		txOpts.Value = value
	}

	return txOpts
}

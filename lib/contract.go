package lib

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/X-mob/mob-watcher/config"
	"github.com/X-mob/mob-watcher/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	tx, err := XmobManageInstance.CreateMob(BasicTransactionOpts, token, tokenId, raisedTotal, takeProfitPrice, stopLossPrice, raisedAmountDeadline, deadline, mobName)
	if err != nil {
		log.Fatalf("create mob failed %s", err)
		panic(err)
	}
	fmt.Printf("tx Hash: %s, wait to be mined..\n", tx.Hash())
	receipt, err := bind.WaitMined(context.Background(), EthClient, tx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("tx mined at %d\n", receipt.BlockNumber.Int64())
}

// only for testing, mainnet we don't have to do this
func PatchMobWithWethSeaport(mobAddress string) {
	mob := GetMobByAddress(mobAddress)
	{
		tx, err := mob.SetSeaportAddress(BasicTransactionOpts, config.GlobalConfig.SeaportAddress)
		if err != nil {
			log.Fatalf("patch mob failed %s", err)
			panic(err)
		}
		fmt.Printf("patch seaport, tx Hash: %s, wait to be mined..\n", tx.Hash())
	}

	{
		tx, err := mob.SetWeth9Address(BasicTransactionOpts, config.GlobalConfig.WethAddress)
		if err != nil {
			log.Fatalf("patch mob failed %s", err)
			panic(err)
		}
		fmt.Printf("patch weth9, tx Hash: %s, wait to be mined..\n", tx.Hash())
	}
}

func JoinMob(value string, address string) {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)
	txOpts.GasPrice = BasicTransactionOpts.GasPrice
	txOpts.Value = utils.StringToBigInt(value)

	mob := GetMobByAddress(address)
	tx, err := mob.JoinPay(txOpts, txOpts.From)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tx Hash: %s, wait to be mined..", tx.Hash())
	receipt, err := bind.WaitMined(context.Background(), EthClient, tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx mined!")
	fmt.Println(receipt)
}

func Claim() {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)
	txOpts.GasPrice = BasicTransactionOpts.GasPrice

	mob := GetMobById(big.NewInt(1))

	// check if claim is allowed
	canClaim, err := mob.CanClaim(nil)
	if err != nil {
		panic(err)
	}
	if canClaim == false {
		panic("can not claim yet")
	}

	// claim
	tx, err := mob.Claim(txOpts)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx Hash is s%, wait for claim mined", tx.Hash())
}

func Settle() {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)
	txOpts.GasPrice = BasicTransactionOpts.GasPrice

	mob := GetMobById(big.NewInt(1))
	tx, err := mob.SettlementAllocation(txOpts, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx Hash is s%, wait for Settle mined", tx.Hash())
}

func Buy() {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)
	txOpts.GasPrice = BasicTransactionOpts.GasPrice

	mob := GetMobById(big.NewInt(1))
	tx, err := mob.SettlementAllocation(txOpts, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx Hash is s%, wait for Settle mined", tx.Hash())
}

func Sell() {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)
	txOpts.GasPrice = BasicTransactionOpts.GasPrice

	mob := GetMobById(big.NewInt(1))
	tx, err := mob.SettlementAllocation(txOpts, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx Hash is s%, wait for Settle mined", tx.Hash())
}

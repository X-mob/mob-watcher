package lib

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/X-mob/mob-watcher/config"
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
	mob, err := NewXmobExchangeCore(mobAddress, EthClient)
	if err != nil {
		panic(err)
	}
	return mob
}

func CreateMob() {
	token := common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")
	tokenId := big.NewInt(1)
	raisedTotal := big.NewInt(100000000)
	takeProfitPrice := big.NewInt(25)
	stopLossPrice := big.NewInt(1)
	raisedAmountDeadline := big.NewInt(time.Now().UnixNano())
	deadline := big.NewInt(time.Now().UnixNano() + 100000)
	mobName := "test mob"

	tx, err := XmobManageInstance.CreateMob(BasicTransactionOpts, token, tokenId, raisedTotal, takeProfitPrice, stopLossPrice, raisedAmountDeadline, deadline, mobName)
	if err != nil {
		log.Fatalf("create mob failed %s", err)
	}
	fmt.Println("tx Hash: s%", tx.Hash())
	// receipt, err := bind.WaitMined(context.Background(), EthClient, tx)
	// if(err != nil){
	// 	panic(err)
	// }
	// for _, log := range receipt.Logs {
	// 	logInterface, err := .Unpack("MobCreate", vLog.Data)
	// 	if err != nil {
	// 	    panic(err)
	// 	}
	// 	fmt.Println(logInterface...);
	// }
}

func JoinMob() {
	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	var txOpts *bind.TransactOpts = bind.NewKeyedTransactor(privateKey)
	txOpts.GasPrice = BasicTransactionOpts.GasPrice
	txOpts.Value = big.NewInt(5)

	mob := GetMobById(big.NewInt(1))
	tx, err := mob.JoinPay(txOpts, txOpts.From)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx Hash is s%, done.", tx.Hash())
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

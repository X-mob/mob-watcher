package lib

import (
	"context"
	"log"
	"math/big"

	"github.com/X-mob/mob-watcher/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	// init client
	url := config.GlobalConfig.RpcUrl
	cli, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	EthClient = cli

	// init contract instance
	instance, err := NewXmobManage(config.GlobalConfig.XmobManagerContractAddress, EthClient)
	if err != nil {
		panic(err)
	}
	XmobManageInstance = instance

	// init BasicKeyTransactor
	privateKey, err := crypto.HexToECDSA(config.GlobalConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	BasicTransactionOpts = bind.NewKeyedTransactor(privateKey)
	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	BasicTransactionOpts.GasPrice = gasPrice
}

var EthClient *ethclient.Client
var XmobManageInstance *XmobManage
var BasicTransactionOpts *bind.TransactOpts

func GetChainId() (*big.Int, error) {
	return EthClient.ChainID(context.Background())
}

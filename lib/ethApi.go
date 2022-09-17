package lib

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/X-mob/mob-watcher/config"
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
	instance, err := NewXmobExchangeCore(config.GlobalConfig.XmobExchangeCoreContractAddress, EthClient)
	if err != nil {
		panic(err)
	}
	XmobExchangeCoreInstance = instance
}

var EthClient *ethclient.Client
var XmobExchangeCoreInstance *XmobExchangeCore

func GetChainId() (*big.Int, error) {
	return EthClient.ChainID(context.Background())
}

func GetVersion() {
	version, err := XmobExchangeCoreInstance.VERSION(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract version:", version)
}

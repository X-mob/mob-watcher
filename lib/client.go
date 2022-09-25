package lib

import (
	"context"
	"math/big"

	"github.com/X-mob/mob-watcher/config"
	"github.com/X-mob/mob-watcher/seaport"
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
	instance, err := NewXmobManage(config.GlobalConfig.XmobManagerAddress, EthClient)
	if err != nil {
		panic(err)
	}
	XmobManageInstance = instance

	// init seaport instance
	seaportInstance, err := seaport.NewSeaport(config.GlobalConfig.SeaportAddress, EthClient)
	if err != nil {
		panic(err)
	}
	SeaportInstance = seaportInstance
}

var EthClient *ethclient.Client
var XmobManageInstance *XmobManage
var SeaportInstance *seaport.Seaport

func GetChainId() (*big.Int, error) {
	return EthClient.ChainID(context.Background())
}

func GetLatestBlockNum() *uint64 {
	latestBlock, err := EthClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}
	return &latestBlock
}

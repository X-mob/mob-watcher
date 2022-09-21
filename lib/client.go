package lib

import (
	"context"
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
	instance, err := NewXmobManage(config.GlobalConfig.XmobManagerContractAddress, EthClient)
	if err != nil {
		panic(err)
	}
	XmobManageInstance = instance
}

var EthClient *ethclient.Client
var XmobManageInstance *XmobManage

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

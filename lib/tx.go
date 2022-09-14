package lib

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	url:="https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"
	cli, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	cli.ChainID(context.Background())
}

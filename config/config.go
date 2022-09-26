package config

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	GlobalConfig = LoadConfig()
}

var GlobalConfig Config

type Config struct {
	RpcUrl              string
	PrivateKey          string
	Network             Network
	OpenSeaApiKey       string
	OpenseaApiV1BaseUrl string
	OpenseaApiV2BaseUrl string
	XmobManagerAddress  common.Address
	SeaportAddress      common.Address
	StorePath           string
	ListenPort          string
}

func LoadConfig() Config {
	network := GetNetwork(GetOptional("NETWORK"))
	fmt.Printf("using network %s..\n", network.String())
	return Config{
		RpcUrl:              GetRequired("RPC_URL"),
		PrivateKey:          GetRequired("PRIVATE_KEY"),
		Network:             network,
		OpenSeaApiKey:       GetRequired("OPENSEA_API_KEY"),
		OpenseaApiV1BaseUrl: GetOpenSeaV1BaseUrl(network),
		OpenseaApiV2BaseUrl: GetOpenSeaV2BaseUrl(network),
		XmobManagerAddress:  common.HexToAddress(GetContractAddress(network).XmobManagerAddress),
		SeaportAddress:      common.HexToAddress(GetContractAddress(network).SeaportAddress),
		StorePath:           GetOptional("STORE_PATH"),
		ListenPort:          GetOptional("LISTEN_PORT"),
	}
}

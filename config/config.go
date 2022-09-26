package config

import (
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
	return Config{
		RpcUrl:              GetRequired("RPC_URL"),
		PrivateKey:          GetRequired("PRIVATE_KEY"),
		Network:             GetNetwork(GetOptional("NETWORK")),
		OpenSeaApiKey:       GetRequired("OPENSEA_API_KEY"),
		OpenseaApiV1BaseUrl: GetOpenSeaV1BaseUrl(GetNetwork(GetOptional("NETWORK"))),
		OpenseaApiV2BaseUrl: GetOpenSeaV2BaseUrl(GetNetwork(GetOptional("NETWORK"))),
		XmobManagerAddress:  common.HexToAddress(GetContractAddress(GetNetwork(GetOptional("NETWORK"))).XmobManagerAddress),
		SeaportAddress:      common.HexToAddress(GetContractAddress(GetNetwork(GetOptional("NETWORK"))).SeaportAddress),
		StorePath:           GetOptional("STORE_PATH"),
		ListenPort:          GetOptional("LISTEN_PORT"),
	}
}

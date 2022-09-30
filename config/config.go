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
	RpcUrl                string
	PrivateKey            string
	Network               Network
	OpenSeaApiKey         string
	OpenseaApiV1BaseUrl   string
	OpenseaApiV2BaseUrl   string
	XmobManagerAddress    common.Address
	SeaportAddress        common.Address
	SeaportConduitAddress common.Address
	SeaportConduitKey     [32]byte
	StorePath             string
	ListenPort            string
}

func LoadConfig() Config {
	rpc := GetRequired("RPC_URL")
	network := GetNetwork(GetOptional("NETWORK"))
	xmobAddr := GetContractAddress(network).XmobManagerAddress
	seaportAddr := GetContractAddress(network).SeaportAddress
	conduitAddr := GetContractAddress(network).SeaportConduitAddress
	fmt.Printf("using network %s\nrpc: %s\nxmobManage: %s\n----\n", network.String(), rpc, xmobAddr)
	return Config{
		RpcUrl:                rpc,
		PrivateKey:            GetRequired("PRIVATE_KEY"),
		Network:               network,
		OpenSeaApiKey:         GetRequired("OPENSEA_API_KEY"),
		OpenseaApiV1BaseUrl:   GetOpenSeaV1BaseUrl(network),
		OpenseaApiV2BaseUrl:   GetOpenSeaV2BaseUrl(network),
		XmobManagerAddress:    common.HexToAddress(xmobAddr),
		SeaportAddress:        common.HexToAddress(seaportAddr),
		SeaportConduitAddress: common.HexToAddress(conduitAddr),
		SeaportConduitKey:     HexStringToByte32(SEAPORT_CONDUIT_KEY),
		StorePath:             GetOptional("STORE_PATH"),
		ListenPort:            GetOptional("LISTEN_PORT"),
	}
}

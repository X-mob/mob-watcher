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

const OpenseaApiV1BaseUrl = "https://api.opensea.io/api/v1"
const OpenseaApiV2BaseUrl = "https://api.opensea.io/v2"

const xmobManagerContractAddress string = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"
const seaportAddress string = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
const wethAddress string = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"

func getRequired(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("[env] Invalid type assertion, not a string")
	}
	return value
}

func getOptional(key string) string {
	value := viper.GetString(key)
	return value
}

type Config struct {
	RpcUrl                     string
	PrivateKey                 string
	XmobManagerContractAddress common.Address
	SeaportAddress             common.Address
	WethAddress                common.Address
	StorePath                  string
	OpenSeaApiKey              string
}

func LoadConfig() Config {
	return Config{
		RpcUrl:                     getRequired("RPC_URL"),
		PrivateKey:                 getRequired("PRIVATE_KEY"),
		OpenSeaApiKey:              getRequired("OPENSEA_API_KEY"),
		XmobManagerContractAddress: common.HexToAddress(xmobManagerContractAddress),
		SeaportAddress:             common.HexToAddress(seaportAddress),
		WethAddress:                common.HexToAddress(wethAddress),
		StorePath:                  getOptional("STORE_PATH"),
	}
}

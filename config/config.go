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

const OpenseaApiV1BaseUrl = "https://testnets-api.opensea.io/api/v1"
const OpenseaApiV2BaseUrl = "https://testnets-api.opensea.io/v2"

const xmobManagerAddress string = "0xe71C9Ec4CE47265d10B01C0cc0994617d43b3AeA"
const seaportAddress string = "0x00000000006c3852cbEf3e08E8dF289169EdE581"

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
	RpcUrl             string
	PrivateKey         string
	XmobManagerAddress common.Address
	SeaportAddress     common.Address
	StorePath          string
	OpenSeaApiKey      string
	ListenPort         string
}

func LoadConfig() Config {
	return Config{
		RpcUrl:             getRequired("RPC_URL"),
		PrivateKey:         getRequired("PRIVATE_KEY"),
		OpenSeaApiKey:      getRequired("OPENSEA_API_KEY"),
		XmobManagerAddress: common.HexToAddress(xmobManagerAddress),
		SeaportAddress:     common.HexToAddress(seaportAddress),
		StorePath:          getOptional("STORE_PATH"),
		ListenPort:         getOptional("LISTEN_PORT"),
	}
}

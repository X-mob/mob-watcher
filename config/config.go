package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"log"
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
var xmobManagerContractAddress string = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"

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
	StorePath                  string
}

func LoadConfig() Config {
	return Config{
		RpcUrl:                     getRequired("RPC_URL"),
		PrivateKey:                 getRequired("PRIVATE_KEY"),
		XmobManagerContractAddress: common.HexToAddress(xmobManagerContractAddress),
		StorePath:                  getOptional("STORE_PATH"),
	}
}

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
var xmobManagerContractAddress string = "0x147B8eb97fD247D06C4006D269c90C1908Fb5D54"
var xmobExchangeCoreContractAddress string = "0x147B8eb97fD247D06C4006D269c90C1908Fb5D54"

func env(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

type Config struct {
	RpcUrl                          string
	PrivateKey                      string
	XmobManagerContractAddress      common.Address
	XmobExchangeCoreContractAddress common.Address
}

func LoadConfig() Config {
	return Config{
		RpcUrl:                          env("RPC_URL"),
		PrivateKey:                      env("PRIVATE_KEY"),
		XmobManagerContractAddress:      common.HexToAddress(xmobManagerContractAddress),
		XmobExchangeCoreContractAddress: common.HexToAddress(xmobExchangeCoreContractAddress),
	}
}

package config

import (
	"encoding/hex"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func GetRequired(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("[env] Invalid type assertion, not a string")
	}
	return value
}

func GetOptional(key string) string {
	value := viper.GetString(key)
	return value
}

func HexStringToByte32(hexString string) [32]byte {
	if hexString == "0x0000000000000000000000000000000000000000000000000000000000000000" {
		return [32]byte{0}
	}

	decodedByteArray, err := hex.DecodeString(strings.TrimPrefix(hexString, "0x"))

	if err != nil {
		panic("Unable to convert hex to byte.")
	}

	var arr [32]byte
	copy(arr[:], decodedByteArray)
	return arr
}

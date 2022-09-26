package config

import (
	"log"

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

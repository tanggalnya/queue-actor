package config

import (
	"log"

	"github.com/spf13/viper"
)

func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Panicf("%s key is not set", key)
	}
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

package config

import "github.com/spf13/viper"

func LoadWorker(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("worker")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			panic(err)
		}
	}

	initGoogleSpreadsheetConfig()

	return nil
}

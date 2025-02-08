package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func MustReadConfig(path string) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("error on load config - %v", err))
	}
}
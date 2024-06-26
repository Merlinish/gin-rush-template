package config

import (
	"fmt"
	"gin-rush-template/tools"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var C = Config{}

func Read(filePath string) {
	viper.SetConfigFile(filePath)
	if tools.FileExist(filePath) {
		tools.PanicOnErr(viper.ReadInConfig())
		tools.PanicOnErr(viper.Unmarshal(&C))
	} else {
		fmt.Println("Config file not exist in ", filePath, ". Using environment variables.")
		tools.PanicOnErr(envconfig.Process("", &C))
	}
}

func Set(config Config) {
	C = config
}

func Get() Config {
	return C
}

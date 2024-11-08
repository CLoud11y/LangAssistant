package config

import (
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	BAIDU_API struct {
		Appid string `mapstructure:"appid"`
		Key   string `mapstructure:"key"`
		Salt  string `mapstructure:"salt"`
	} `mapstructure:"BAIDU_API"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic("Unable to decode into struct: " + err.Error())
	}
}

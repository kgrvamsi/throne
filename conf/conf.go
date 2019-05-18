package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Default DefaultInfo
	Log     LogInfo
}

type DefaultInfo struct {
	Port string `toml:"port"`
}

type LogInfo struct {
	Type     string `toml:"type"`
	LogLevel string `toml:"log_level"`
}

//GetConf ...Fetch the configuration from the config.toml file
func GetConf() Config {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Errorf(err.Error())
	}
	return conf
}

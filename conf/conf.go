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
	LogType     string `toml:"log_type"`
	LogLevel    string `toml:"log_level"`
	ESUrl       string `toml:"es_url"`
	ESHttps     bool   `toml:"es_https"`
	ESPort      string `toml:"es_port"`
	ESUsername  string `toml:"es_username"`
	ESPassword  string `toml:"es_password"`
	ESIndexName string `toml:"es_index_name"`
}

//GetConf ...Fetch the configuration from the config.toml file
func GetConf() Config {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Errorf(err.Error())
	}
	return conf
}

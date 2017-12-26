package crud

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Conf hold database configuration stuff
type Conf struct {
	DbUser string
	DbPass string
	DbName string
}

//GetConf gets a conf!
func GetConf() Conf {
	var conf Conf
	if _, err := toml.DecodeFile("crud/conf/config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	return conf
}

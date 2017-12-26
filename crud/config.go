package engine

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	dbUser string
	dbPass string
	dbName string
}

func (c *conf) getConf() *conf {

	ymlFile, err := ioutil.ReadFile("conf/config.yml")
	if err != nil {
		log.Printf("ymlfile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

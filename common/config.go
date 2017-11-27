package common

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	ServerAddress string `json:"server_address"`
	MongoAddress  string `json:"mongo_address"`
}

var C config

func init() {
	InitConfig()
}

func InitConfig() {
	f, err := os.Open("common/config.json")
	if err != nil {
		log.Fatalln("Can't open config file", err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&C)
	if err != nil {
		log.Fatalln("Can't decode config json", err)
	}
}

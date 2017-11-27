package common

import (
	"encoding/json"
	"log"
	"os"
	"pool/logger"
)

type config struct {
	ServerAddress string    `json:"server_address"`
	MongoAddress  string    `json:"mongo_address"`
	Logger        logConfig `json:"log_config"`
}

type logConfig struct {
	filePath   string `json:"file_path"`
	maxSize    int    `json:"max_size_MB"`
	maxAge     int    `json:"max_age_days"`
	maxBackups int    `json:"max_backups"`
}

var (
	C config
)

func initConfig() {
	f, err := os.Open("common/config.json")
	if err != nil {
		log.Fatalln("Can't open config file", err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&C)
	if err != nil {
		log.Fatalf("Can't decode config -> %s", err)
	}
}

func initLogger() {
	l := logger.NewLogger(C.Logger.filePath, C.Logger.maxSize, C.Logger.maxAge, C.Logger.maxBackups)
	log.SetOutput(l)
}

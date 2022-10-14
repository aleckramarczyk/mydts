package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	API_Host     string `mapstructure:"API_Host"`
	API_Endpoint string `mapstructure:"API_Endpoint"`
}

var AgentConfig *Config

func LoadConfig() {
	log.Println("Attempting to load agent configuration...")

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(exPath)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AgentConfig)
	if err != nil {
		log.Println(err)
	}
	echoConf()
}

func echoConf() {
	log.Printf("Config loaded. API_Host: %s API_Endpoint: %s", AgentConfig.API_Host, AgentConfig.API_Endpoint)
}

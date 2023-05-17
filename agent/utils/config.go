package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	AgentID                 string        `mapstructure:"SerialNumber"`
	ApiHost                 string        `mapstructure:"ApiHost"`
	ApiEndpoint             string        `mapstructure:"ApiEndpoint"`
	ApplicationStatePath    string        `mapstructure:"ApplicationStatePath"`
	FromClientGPSPath       string        `mapstructure:"FromClientGPSPath"`
	FromClientAuthorizePath string        `mapstructure:"FromClientAuthorizePath"`
	FromClientSignOffPath   string        `mapstructure:"FromClientSignOffPath"`
	UpdateInterval          time.Duration `mapstructure:"UpdateInterval"`
	UpdateOnInterval        bool          `mapstructure:"UpdateOnInterval"`
}

var AgentConfig *Config

func LoadConfig() {
	log.Println("Attempting to load agent configuration...")

	// Gets the location where the executable is running, this is used to determine the path to where the config file
	// should be
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	viper.AddConfigPath(exPath)

	/*
		// Read configuration information that is persistent after upgrades. Stored in persistentValues.json
		viper.SetConfigName("persistentValues")
		viper.SetConfigType("json")
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalf("ERROR reading persistentValues: %s", err)
		}
	*/

	// Read configuration data that can change between upgrades. Stored in config.json
	viper.SetConfigName("config")
	err = viper.MergeInConfig()
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
	log.Printf("Config loaded. API_Host: %s API_Endpoint: %s", AgentConfig.ApiHost, AgentConfig.ApiEndpoint)
}

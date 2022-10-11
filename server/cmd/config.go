package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_port     string `mapstructure:"DB_PORT"`
	DB_host     string `mapstructure:"DB_HOST"`
	DB_table    string `mapstructure:"DB_TABLE"`
	DB_user     string `mapstructure:"DB_USER"`
	DB_password string `mapstructure:"DB_PASSWORD"`
	API_port    string `mapstructure:"API_PORT"`
	DEBUG       int    `mapstructure:"DEBUG"`
}

var AppConfig *Config

func LoadAppConfig(path string) {
	log.Println("Loading server configurations...")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_TABLE")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("API_PORT")
	viper.BindEnv("DEBUG")
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		log.Println(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Println(err)
	}
}

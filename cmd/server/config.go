package main

import "log"

type Config struct {
	DB_port  string `mapstructure:"db_port"`
	DB_host  string `mapstructure:"db_host"`
	DB_table string `mapstructure:"db_table"`
	DB_user  string `mapstructure:"db_user"`
	API_port string `mapstructure:"api_port"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading server configurations...")
}

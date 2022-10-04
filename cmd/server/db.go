package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectDB(connectionString string) {
	log.Println("Attempting to connect to database...")
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}
	log.Println("Connected to database")
}

func GenerateConfigString() string {
	configString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", AppConfig.DB_user, AppConfig.DB_password, AppConfig.DB_host, AppConfig.DB_port, AppConfig.DB_table)
	return configString
}

func MDTExists(id string) bool {
	var mdt MDT
	Instance.First(&mdt, id)
	return mdt.ID != ""
}

func Migrate() {
	Instance.AutoMigrate(&MDT{})
	log.Println("Database migration completed")
}

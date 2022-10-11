package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectDB(connectionString string) error {
	log.Println("Attempting to connect to database...")
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connected to database")
	return nil
}

func GenerateConfigString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", AppConfig.DB_user, AppConfig.DB_password, AppConfig.DB_host, AppConfig.DB_port, AppConfig.DB_table)
}

/*
func MDTExists(mac string) bool {
	var mdt MDT
	Instance.First(&mdt, mac)
	if mdt.Dock_mac == "" {
		return false
	}
	return true
}
*/
func MDTExists(mac string) bool {
	var mdt MDT
	r := Instance.Where("dock_mac = ?", mac).Limit(1).Find(&mdt)
	return r.RowsAffected > 0
}

func Migrate() error {
	err := Instance.AutoMigrate(&MDT{})
	if err != nil {
		return err
	}
	log.Println("Database migration completed")
	return nil
}

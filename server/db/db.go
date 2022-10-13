package db

import (
	"fmt"
	"log"

	"aleckramarczyk/mydts/server/entities"
	"aleckramarczyk/mydts/server/utils"

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
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.AppConfig.DB_user, utils.AppConfig.DB_password, utils.AppConfig.DB_host, utils.AppConfig.DB_port, utils.AppConfig.DB_table)
}

func MDTExists(mac string) bool {
	var mdt entities.MDT
	r := Instance.Where("dock_mac = ?", mac).Limit(1).Find(&mdt)
	return r.RowsAffected > 0
}

func Migrate() error {
	err := Instance.AutoMigrate(&entities.MDT{})
	if err != nil {
		return err
	}
	log.Println("Database migration completed")
	return nil
}

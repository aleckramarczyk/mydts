package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	Units map[string]*Unit
)

type Unit struct {
	AgentID      string    `json:"agentID" gorm:"primary_key"`
	UnitName     string    `json:"unitName"`
	UnitId       string    `json:"unitId"`
	VehicleID    string    `json:"vehicleID"`
	InternalIP   string    `json:"internalIP"`
	RemoteIP     string    `json:"remoteIP"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	GpsTimestamp string    `json:"gpsTimestamp"`
	LastUpdated  time.Time `json:"lastUpdated"`
}

var DB *gorm.DB

func init() {
	var err error
	//TODO this is for testing only, this configuration should be able to be configured
	dsn := "host=localhost user=gorm password=gorm dbname=units sslmode=disable TimeZone=America/Phoenix"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fatal error, unable to connect to database:", err)
	}

	err = DB.AutoMigrate(&Unit{})
	if err != nil {
		log.Fatal(err)
	}
}

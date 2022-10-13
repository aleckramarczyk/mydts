package entities

import "time"

type MDT struct {
	Dock_mac  string    `gorm:"primaryKey" json:"dock_mac"`
	Mdt_uuid  string    `gorm:"not null" json:"mdt_uuid"`
	Remote_ip string    `gorm:"not null" json:"ip"`
	Updated   time.Time `gorm:"not null" json:"updated"`
}

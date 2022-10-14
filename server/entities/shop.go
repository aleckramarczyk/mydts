package entities

type Shop struct {
	Dock_mac    string `gorm:"primaryKey" json:"dock_mac"`
	Shop_number string `gorm:"not null" json:"shop_number"`
}

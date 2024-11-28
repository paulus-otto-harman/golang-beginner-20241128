package model

type Order struct { //user
	ID            uint          `gorm:"primaryKey"`
	TransactionID string        `gorm:"size:30;unique" json:"transaction_id" binding:"required"`
	CourierID     uint          `gorm:"not null" json:"courier_id"`
	Courier       Courier       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"courier"`
	Origin        Coordinate    `gorm:"embedded;embeddedPrefix:origin_" json:"origin"`
	Destination   Coordinate    `gorm:"embedded;embeddedPrefix:destination_" json:"destination"`
	Total         float64       `gorm:"total" json:"total" binding:"required"`
	OrderStatus   []OrderStatus `gorm:"foreignKey:OrderID" json:"-"`
	Status        string        `json:"status"`
}

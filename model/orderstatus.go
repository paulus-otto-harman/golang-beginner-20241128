package model

type OrderStatus struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	OrderID  uint   `gorm:"not null" json:"-"`
	Status   string `gorm:"not null, size=10" json:"status"`
	Location string `gorm:"not null" json:"location"`
	//Order   Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

package model

type Courier struct { // company
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `gorm:"size:15;unique" json:"name"`
	Orders []Order `gorm:"foreignKey:CourierID" json:"orders,omitempty"`
}

func CourierSeed() []Courier {
	return []Courier{
		{Name: "JNE"},
		{Name: "TIKI"},
		{Name: "POS"},
	}
}

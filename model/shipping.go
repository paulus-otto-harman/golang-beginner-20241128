package model

type Shipping struct {
	Items       int        `json:"items"`
	CourierId   int        `json:"courier_id"`
	Destination Coordinate `json:"destination"`
	Cost        float32    `json:"cost"`
}

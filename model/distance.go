package model

type Distance struct {
	Code   string `json:"code"`
	Routes []struct {
		Legs []struct {
			Steps    []interface{} `json:"steps"`
			Summary  string        `json:"summary"`
			Weight   float64       `json:"weight"`
			Duration float64       `json:"duration"`
			Distance float64       `json:"distance"`
		} `json:"legs"`
		WeightName string  `json:"weight_name"`
		Weight     float64 `json:"weight"`
		Duration   float64 `json:"duration"`
		Distance   float64 `json:"distance"`
	} `json:"routes"`
	Waypoints []struct {
		Hint     string    `json:"hint"`
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
}

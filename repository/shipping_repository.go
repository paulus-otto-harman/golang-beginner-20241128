package repository

import (
	"20241128/model"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type ShippingRepository interface {
	Cost(shipping *model.Shipping) error
}

type shippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) ShippingRepository {
	return &shippingRepository{db: db}
}

func (r *shippingRepository) Cost(shipping *model.Shipping) error {
	origin := viper.GetString("STORE_LOCATION")
	client := &http.Client{}

	geoApiBaseUrl := "https://router.project-osrm.org"
	geoApiServiceUrl := "route/v1/driving"
	destination := fmt.Sprintf("%v,%v", shipping.Destination.Lat, shipping.Destination.Lng)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s;%s", geoApiBaseUrl, geoApiServiceUrl, origin, destination), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	var distance model.Distance
	json.Unmarshal(body, &distance)

	shipping.Cost = float32(shipping.Items) * float32(distance.Routes[0].Distance) * 2000

	return nil
}

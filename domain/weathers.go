package domain

import "time"

type Weathers []Weather

type Weather struct {
	SoilTemperature *int      `json:"soilTemperature"`
	SoilMoisture    *int      `json:"soilMoisture"`
	WindSpeed       *int      `json:"windSpeed"`
	WindDirection   string    `json:"windDirection"`
	Temperature     *int      `json:"temperature"`
	Pressure        *int      `json:"pressure"`
	Humidity        *int      `json:"humidity"`
	Altitude        *int      `json:"altitude"`
	Rain            *float64  `json:"rain"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

package domain

import "time"

type Weathers []Weather

type Weather struct {
	Id   string    `json:"id"`
	Time time.Time `json:"time"`
	Lati float64   `json:"lati,string"`
	Long float64   `json:"long,string"`
	Temp float64   `json:"temp,string"`
	Humi float64   `json:"humi,string"`
	Baro float64   `json:"baro,string"`
	Rain float64   `json:"rain,string"`
	Wind struct {
		Spee float64 `json:"spee,string"`
		Dire float64 `json:"dire,string"`
		Gust float64 `json:"gust,string"`
	} `json:"wind"`
	Ligh struct {
		Dura float64 `json:"dura,string"`
		Amou float64 `json:"amou,string"`
	} `json:"ligh"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

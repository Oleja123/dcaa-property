package property

import "time"

type Property struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Addr       string    `json:"addr"`
	Price      float64   `json:"price"`
	Info       string    `json:"info"`
	CategoryId string    `json:"category_id"`
	LastUpdate time.Time `json:"last_update"`
}

package propertydto

import categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"

type PropertyDTO struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Addr       string  `json:"addr"`
	Price      float64 `json:"price"`
	Info       string  `json:"info"`
	CategoryId int     `json:"category_id"`
	LastUpdate string  `json:"last_update"`
}

type ExtendedPropertyDTO struct {
	Id         int                     `json:"id"`
	Name       string                  `json:"name"`
	Addr       string                  `json:"addr"`
	Price      float64                 `json:"price"`
	Info       string                  `json:"info"`
	CategoryId int                     `json:"category_id"`
	Category   categorydto.CategoryDTO `json:"category"`
	LastUpdate string                  `json:"last_update"`
}

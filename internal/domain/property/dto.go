package property

import "github.com/Oleja123/dcaa-property/internal/domain/category"

type PropertyDTO struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Addr       string  `json:"addr"`
	Price      float64 `json:"price"`
	Info       string  `json:"info"`
	CategoryId string  `json:"category_id"`
}

type ExtendedPropertyDTO struct {
	Id         int                  `json:"id"`
	Name       string               `json:"name"`
	Addr       string               `json:"addr"`
	Price      float64              `json:"price"`
	Info       string               `json:"info"`
	CategoryId string               `json:"category_id"`
	Category   category.CategoryDTO `json:"category"`
}

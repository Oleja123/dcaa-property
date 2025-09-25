package propertydto

import (
	categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"
	optionalType "github.com/denpa16/optional-go-type"
)

type PropertyDTO struct {
	Id         optionalType.OptionalInt     `json:"id"`
	Name       optionalType.OptionalString  `json:"name"`
	Addr       optionalType.OptionalString  `json:"addr"`
	Price      optionalType.OptionalFloat64 `json:"price"`
	Info       optionalType.OptionalString  `json:"info"`
	CategoryId optionalType.OptionalInt     `json:"category_id"`
	LastUpdate optionalType.OptionalString  `json:"last_update"`
}

func (p *PropertyDTO) Validate(isUpdate bool) bool {
	if !p.Name.Valid || !p.Addr.Valid || !p.CategoryId.Valid || (isUpdate && !p.Id.Valid) ||
		!p.Price.Defined || !p.Info.Defined {
		return false
	}
	return true
}

type ExtendedPropertyDTO struct {
	PropertyDTO
	Category categorydto.CategoryDTO `json:"category"`
}

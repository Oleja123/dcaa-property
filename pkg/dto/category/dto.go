package categorydto

import optionalType "github.com/denpa16/optional-go-type"

type CategoryDTO struct {
	Id   optionalType.OptionalInt    `json:"id"`
	Name optionalType.OptionalString `json:"name"`
	Info optionalType.OptionalString `json:"info"`
}

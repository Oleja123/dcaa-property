package categorydto

import optionalType "github.com/denpa16/optional-go-type"

type CategoryDTO struct {
	Id   optionalType.OptionalInt    `json:"id"`
	Name optionalType.OptionalString `json:"name"`
	Info optionalType.OptionalString `json:"info"`
}

func (c *CategoryDTO) Validate(isUpdate bool) bool {
	if !c.Name.Valid || (isUpdate && !c.Id.Valid) || !c.Info.Defined {
		return false
	}
	return true
}

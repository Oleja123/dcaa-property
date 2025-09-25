package category

import (
	categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"
)

type Service interface {
	FindOne(id int) (categorydto.CategoryDTO, error)
}

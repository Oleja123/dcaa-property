package categorymock

import (
	"errors"

	categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"
	optionalType "github.com/denpa16/optional-go-type"
)

type MockCategoryClient struct{}

func (m *MockCategoryClient) FindOne(id int) (categorydto.CategoryDTO, error) {
	if id == 0 {
		return categorydto.CategoryDTO{}, errors.New("not found")
	}
	name := "Category"
	info := "Info"
	return categorydto.CategoryDTO{
		Id:   optionalType.NewOptionalInt(&id),
		Name: optionalType.NewOptionalString(&name),
		Info: optionalType.NewOptionalString(&info),
	}, nil
}

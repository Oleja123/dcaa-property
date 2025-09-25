package property

import (
	"context"

	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
)

type Service interface {
	Create(ctx context.Context, dto propertydto.PropertyDTO) (int, error)
	FindAll(ctx context.Context) ([]propertydto.PropertyDTO, error)
	FindOne(ctx context.Context, id int) (propertydto.ExtendedPropertyDTO, error)
	Update(ctx context.Context, property propertydto.PropertyDTO) error
	Delete(ctx context.Context, id int) error
}

package property

import "context"

type Service interface {
	Create(ctx context.Context, dto PropertyDTO) (int, error)
	FindAll(ctx context.Context) ([]PropertyDTO, error)
	FindOne(ctx context.Context, id int) (PropertyDTO, error)
	Update(ctx context.Context, property PropertyDTO) error
	Delete(ctx context.Context, id int) error
}

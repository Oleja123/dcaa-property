package property

import "context"

type Repository interface {
	Create(ctx context.Context, property Property) (int, error)
	FindAll(ctx context.Context) ([]Property, error)
	FindOne(ctx context.Context, id int) (Property, error)
	Update(ctx context.Context, property Property) error
	Delete(ctx context.Context, id int) error
}

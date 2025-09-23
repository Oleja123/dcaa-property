package propertymock

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
)

type MockPropertyRepo struct{}

func (m *MockPropertyRepo) Create(ctx context.Context, p property.Property) (int, error) {
	if p.Name == "fail" {
		return 0, errors.New("creation failed")
	}
	return 1, nil
}

func (m *MockPropertyRepo) Update(ctx context.Context, p property.Property) error {
	if p.Id == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (m *MockPropertyRepo) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (m *MockPropertyRepo) FindAll(ctx context.Context) ([]property.Property, error) {
	return []property.Property{
		{
			Id:         1,
			Name:       "House",
			Addr:       "Main St",
			Price:      sql.NullFloat64{Float64: 100000, Valid: true},
			Info:       sql.NullString{String: "Nice house", Valid: true},
			CategoryId: 2,
			LastUpdate: time.Now(),
		},
	}, nil
}

func (m *MockPropertyRepo) FindOne(ctx context.Context, id int) (property.Property, error) {
	if id == 0 {
		return property.Property{}, errors.New("not found")
	}
	return property.Property{
		Id:         id,
		Name:       "House",
		Addr:       "Main St",
		Price:      sql.NullFloat64{Float64: 100000, Valid: true},
		Info:       sql.NullString{String: "Nice house", Valid: true},
		CategoryId: 2,
		LastUpdate: time.Now(),
	}, nil
}

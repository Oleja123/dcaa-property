package propertyservice_test

import (
	"context"
	"testing"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertymock "github.com/Oleja123/dcaa-property/internal/repository/property"
	propertyservice "github.com/Oleja123/dcaa-property/internal/service/property"
	"github.com/stretchr/testify/assert"
)

func TestPropertyService(t *testing.T) {
	ctx := context.Background()
	repo := &propertymock.MockPropertyRepo{}
	service := propertyservice.NewService(repo)

	dto := property.PropertyDTO{
		Name:       "Villa",
		Addr:       "Ocean Drive",
		Price:      500000,
		Info:       "Sea view",
		CategoryId: 1,
	}
	id, err := service.Create(ctx, dto)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	dto.Id = 1
	dto.Price = 600000
	err = service.Update(ctx, dto)
	assert.NoError(t, err)

	err = service.Delete(ctx, 1)
	assert.NoError(t, err)

	all, err := service.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "House", all[0].Name)

	found, err := service.FindOne(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "House", found.Name)
}

package propertyservice_test

import (
	"context"
	"testing"

	categoryservice "github.com/Oleja123/dcaa-property/internal/application/category"
	propertyservice "github.com/Oleja123/dcaa-property/internal/application/property"
	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
	categorymock "github.com/Oleja123/dcaa-property/internal/infrastructure/category"
	propertymock "github.com/Oleja123/dcaa-property/internal/infrastructure/property"
	myErrors "github.com/Oleja123/dcaa-property/pkg/errors"
	optionalType "github.com/denpa16/optional-go-type"
	"github.com/stretchr/testify/assert"
)

func TestPropertyService(t *testing.T) {
	ctx := context.Background()
	repo := &propertymock.MockPropertyRepo{}
	client := &categorymock.MockCategoryClient{}
	se := categoryservice.NewService(client)
	service := propertyservice.NewService(repo, se)

	name := "Villa"
	addr := "Ocean Drive"
	price := 500000.0
	info := "Sea view"
	categoryId := 1

	dto := propertydto.PropertyDTO{
		Name:       optionalType.NewOptionalString(&name),
		Addr:       optionalType.NewOptionalString(&addr),
		Price:      optionalType.NewOptionalFloat64(&price),
		Info:       optionalType.NewOptionalString(&info),
		CategoryId: optionalType.NewOptionalInt(&categoryId),
	}
	id, err := service.Create(ctx, dto)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	categoryId = 0
	dto.CategoryId = optionalType.NewOptionalInt(&categoryId)
	id, err = service.Create(ctx, dto)
	assert.ErrorAs(t, err, &myErrors.ErrNotFound)
	assert.Equal(t, 0, id)

	categoryId = 1
	id = 1
	dto.CategoryId = optionalType.NewOptionalInt(&categoryId)
	dto.Id = optionalType.NewOptionalInt(&id)
	*dto.Price.Value = 600000
	err = service.Update(ctx, dto)
	assert.NoError(t, err)

	categoryId = 0
	dto.CategoryId = optionalType.NewOptionalInt(&categoryId)
	dto.Id = optionalType.NewOptionalInt(&id)
	*dto.Price.Value = 600000
	err = service.Update(ctx, dto)
	assert.ErrorAs(t, err, &myErrors.ErrNotFound)

	categoryId = 1
	dto.CategoryId = optionalType.NewOptionalInt(&categoryId)

	err = service.Delete(ctx, 1)
	assert.NoError(t, err)

	err = service.Delete(ctx, 0)
	assert.ErrorAs(t, err, &myErrors.ErrNotFound)

	all, err := service.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "House", *all[0].Name.Value)

	found, err := service.FindOne(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "House", *found.Name.Value)

	found, err = service.FindOne(ctx, 0)
	assert.ErrorAs(t, err, &myErrors.ErrNotFound)
}

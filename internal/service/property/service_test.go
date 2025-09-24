package propertyservice_test

import (
	"context"
	"testing"

	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
	propertymock "github.com/Oleja123/dcaa-property/internal/repository/property"
	propertyservice "github.com/Oleja123/dcaa-property/internal/service/property"
	optionalType "github.com/denpa16/optional-go-type"
	"github.com/stretchr/testify/assert"
)

func TestPropertyService(t *testing.T) {
	ctx := context.Background()
	repo := &propertymock.MockPropertyRepo{}
	service := propertyservice.NewService(repo)

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

	dto.Id = optionalType.NewOptionalInt(&id)
	*dto.Price.Value = 600000
	err = service.Update(ctx, dto)
	assert.NoError(t, err)

	err = service.Delete(ctx, 1)
	assert.NoError(t, err)

	all, err := service.FindAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "House", *all[0].Name.Value)

	found, err := service.FindOne(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "House", *found.Name.Value)
}

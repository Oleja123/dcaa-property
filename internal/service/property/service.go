package propertyservice

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
	optionalType "github.com/denpa16/optional-go-type"
)

type propertyService struct {
	repository property.Repository
}

func (ps *propertyService) PropertyToDTO(p property.Property) propertydto.PropertyDTO {
	dto := propertydto.PropertyDTO{}
	dto.Addr = optionalType.NewOptionalString(&p.Addr)
	dto.CategoryId = optionalType.NewOptionalInt(&p.CategoryId)
	dto.Name = optionalType.NewOptionalString(&p.Name)
	if !p.Price.Valid {
		dto.Price = optionalType.NewOptionalFloat64(nil)
	} else {
		dto.Price = optionalType.NewOptionalFloat64(&p.Price.Float64)
	}
	if !p.Info.Valid {
		dto.Info = optionalType.NewOptionalString(nil)
	} else {
		dto.Info = optionalType.NewOptionalString(&p.Info.String)
	}
	dto.Id = optionalType.NewOptionalInt(&p.Id)
	lastUpdateStr := p.LastUpdate.Format(time.RFC3339)
	dto.LastUpdate = optionalType.NewOptionalString(&lastUpdateStr)
	return dto
}

func (ps *propertyService) PropertyFromDTO(ctx context.Context, dto propertydto.PropertyDTO) property.Property {
	p := property.Property{}
	p.Addr = *dto.Addr.Value
	p.CategoryId = *dto.CategoryId.Value
	p.Name = *dto.Name.Value
	if dto.Id.Valid {
		p.Id = *dto.Id.Value
	}
	if dto.Price.Valid {
		p.Price = sql.NullFloat64{
			Float64: *dto.Price.Value,
			Valid:   true,
		}
	}
	if dto.Info.Valid {
		p.Info = sql.NullString{
			String: *dto.Info.Value,
			Valid:  true,
		}
	}
	return p
}

func (ps *propertyService) Create(ctx context.Context, dto propertydto.PropertyDTO) (int, error) {
	property := ps.PropertyFromDTO(ctx, dto)
	id, err := ps.repository.Create(ctx, property)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("ошибка при создании сущности собственности")
	}
	return id, nil
}

func (ps *propertyService) Update(ctx context.Context, dto propertydto.PropertyDTO) error {
	pr := ps.PropertyFromDTO(ctx, dto)
	err := ps.repository.Update(ctx, pr)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("ошибка при обновлении сущности собственности с id: %d", *dto.Id.Value)
	}
	return nil
}

func (ps *propertyService) Delete(ctx context.Context, id int) error {
	err := ps.repository.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("ошибка при удалении сущности собственности с id: %d", id)
	}
	return nil
}

func (ps *propertyService) FindAll(ctx context.Context) ([]propertydto.PropertyDTO, error) {
	pr, err := ps.repository.FindAll(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("ошибка при получении списка записей собственностей")
	}
	res := make([]propertydto.PropertyDTO, 0, len(pr))
	for _, val := range pr {
		res = append(res, ps.PropertyToDTO(val))
	}

	return res, nil
}

func (ps *propertyService) FindOne(ctx context.Context, id int) (propertydto.PropertyDTO, error) {
	pr, err := ps.repository.FindOne(ctx, id)
	if err != nil {
		fmt.Println(err)
		return propertydto.PropertyDTO{}, fmt.Errorf("ошибка при получении записи собственности с id: %d", id)
	}

	return ps.PropertyToDTO(pr), nil
}

func NewService(repo property.Repository) property.Service {
	return &propertyService{
		repository: repo,
	}
}

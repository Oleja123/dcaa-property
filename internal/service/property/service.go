package propertyservice

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
)

type propertyService struct {
	repository property.Repository
}

func (ps *propertyService) PropertyToDTO(p property.Property) propertydto.PropertyDTO {
	dto := propertydto.PropertyDTO{}
	dto.Addr = p.Addr
	dto.CategoryId = p.CategoryId
	dto.Name = p.Name
	dto.Price = p.Price.Float64
	dto.Info = p.Info.String
	dto.Id = p.Id
	dto.LastUpdate = p.LastUpdate.Format(time.RFC3339)
	return dto
}

func (ps *propertyService) PropertyFromDTO(ctx context.Context, dto propertydto.PropertyDTO) property.Property {
	p := property.Property{}
	if dto.Id != 0 {
		if res, err := ps.repository.FindOne(ctx, dto.Id); err == nil {
			p = res
		}
		p.Id = dto.Id
	}
	if dto.Addr != "" {
		p.Addr = dto.Addr
	}
	if dto.CategoryId != 0 {
		p.CategoryId = dto.CategoryId
	}
	if dto.Name != "" {
		p.Name = dto.Name
	}
	if dto.Price != 0 {
		p.Price = sql.NullFloat64{
			Float64: dto.Price,
			Valid:   true,
		}
	}
	if dto.Info != "" {
		p.Info = sql.NullString{
			String: dto.Info,
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
		return fmt.Errorf("ошибка при обновлении сущности собственности с id: %d", dto.Id)
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

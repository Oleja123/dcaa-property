package propertyservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Oleja123/dcaa-property/internal/domain/category"
	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
	categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"
	myErrors "github.com/Oleja123/dcaa-property/pkg/errors"
	optionalType "github.com/denpa16/optional-go-type"
)

type Service struct {
	repository      property.Repository
	categoryService category.Service
}

func (ps *Service) PropertyToDTO(p property.Property) propertydto.PropertyDTO {
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

func (ps *Service) PropertyDTOToExtended(p propertydto.PropertyDTO, ca categorydto.CategoryDTO) propertydto.ExtendedPropertyDTO {
	return propertydto.ExtendedPropertyDTO{PropertyDTO: p, Category: ca}
}

func (ps *Service) PropertyFromDTO(ctx context.Context, dto propertydto.PropertyDTO) property.Property {
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

func (ps *Service) Create(ctx context.Context, dto propertydto.PropertyDTO) (int, error) {
	pr := ps.PropertyFromDTO(ctx, dto)
	if _, err := ps.categoryService.FindOne(pr.CategoryId); err != nil {
		return 0, fmt.Errorf("ошибка при создании сущности собственности: %w", err)
	}
	id, err := ps.repository.Create(ctx, pr)
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("ошибка при создании сущности собственности: %w", err)
	}
	return id, nil
}

func (ps *Service) Update(ctx context.Context, dto propertydto.PropertyDTO) error {
	pr := ps.PropertyFromDTO(ctx, dto)
	if _, err := ps.repository.FindOne(ctx, pr.Id); err != nil {
		return fmt.Errorf("ошибка при обновлении сущности с id: %d: %w", pr.Id, err)
	}
	if _, err := ps.categoryService.FindOne(pr.CategoryId); err != nil {
		return fmt.Errorf("ошибка при обновлении сущности с id: %d: %w", pr.Id, err)
	}
	err := ps.repository.Update(ctx, pr)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("ошибка при обновлении сущности собственности с id: %d: %w", *dto.Id.Value, err)
	}
	return nil
}

func (ps *Service) Delete(ctx context.Context, id int) error {
	if _, err := ps.repository.FindOne(ctx, id); err != nil {
		return fmt.Errorf("ошибка при удалении сущности с id: %d: %w", id, err)
	}
	err := ps.repository.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("ошибка при удалении сущности собственности с id: %d: %w", id, err)
	}
	return nil
}

func (ps *Service) FindAll(ctx context.Context) ([]propertydto.PropertyDTO, error) {
	pr, err := ps.repository.FindAll(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("ошибка при получении списка записей собственностей: %w", err)
	}
	res := make([]propertydto.PropertyDTO, 0, len(pr))
	for _, val := range pr {
		res = append(res, ps.PropertyToDTO(val))
	}

	return res, nil
}

func (ps *Service) FindOne(ctx context.Context, id int) (propertydto.ExtendedPropertyDTO, error) {
	pr, err := ps.repository.FindOne(ctx, id)
	switch {
	case errors.Is(err, myErrors.ErrNotFound):
		return propertydto.ExtendedPropertyDTO{}, fmt.Errorf("ошибка при получении записи собственности с id: %d: %w", id, err)
	case err != nil:
		return propertydto.ExtendedPropertyDTO{}, fmt.Errorf("ошибка при получении записи собственности с id: %d: %w", id, err)
	}
	ca, err := ps.categoryService.FindOne(pr.CategoryId)
	switch {
	case errors.Is(err, myErrors.ErrNotFound):
		ps.Delete(ctx, id)
		return propertydto.ExtendedPropertyDTO{}, fmt.Errorf("ошибка при получении записи собственности с id: %d: %w", id, err)
	case err != nil:
		return propertydto.ExtendedPropertyDTO{}, fmt.Errorf("ошибка при получении записи собственности с id: %d: %w", id, err)
	}

	return ps.PropertyDTOToExtended(ps.PropertyToDTO(pr), ca), nil
}

func NewService(repo property.Repository, cs category.Service) *Service {
	return &Service{
		repository:      repo,
		categoryService: cs,
	}
}

package categoryservice

import (
	"fmt"

	"github.com/Oleja123/dcaa-property/internal/domain/category"
	categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"
)

type Service struct {
	client category.Client
}

func (s *Service) FindOne(id int) (categorydto.CategoryDTO, error) {
	res, err := s.client.FindOne(id)
	if err != nil {
		return categorydto.CategoryDTO{}, fmt.Errorf("не найдена категория с id: %d: %w", id, err)
	}
	return res, nil
}

func NewService(cl category.Client) *Service {
	return &Service{
		client: cl,
	}
}

package service

import (
	"context"
	dto "dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"

	"github.com/mashingan/smapping"
)

type CityService interface {
	CreateCity(ctx context.Context, city dto.AddCityDTO) (entity.City, error)
	GetAllCity(ctx context.Context) ([]entity.City, error)
	GetCityByID(ctx context.Context, id string) (entity.City, error)
	DeleteCity(ctx context.Context, id string) error
	UpdateCity(ctx context.Context, id string, city dto.UpdateCityDTO) (entity.City, error)
}

type cityService struct {
	cityRepository repository.CityRepository
}

func NewCityService(repo repository.CityRepository) CityService {
	return &cityService{
		cityRepository: repo,
	}
}

func (c *cityService) CreateCity(ctx context.Context, city dto.AddCityDTO) (entity.City, error) {
	var createdCity entity.City

	err := smapping.FillStruct(&createdCity, smapping.MapFields(&city))
	if err != nil {
		return entity.City{}, err
	}

	result, err := c.cityRepository.InsertCity(ctx, createdCity)

	return result, err
}

func (c *cityService) GetAllCity(ctx context.Context) ([]entity.City, error) {
	result, err := c.cityRepository.GetAllCity(ctx)

	return result, err
}

func (c *cityService) GetCityByID(ctx context.Context, id string) (entity.City, error) {
	result, err := c.cityRepository.GetCityByID(ctx, id)

	return result, err
}

func (c *cityService) UpdateCity(ctx context.Context, id string, city dto.UpdateCityDTO) (entity.City, error) {
	var updatedCity entity.City

	err := smapping.FillStruct(&updatedCity, smapping.MapFields(&city))
	if err != nil {
		return entity.City{}, err
	}

	result, err := c.cityRepository.UpdateCityByID(ctx, id, updatedCity)

	return result, err
}

func (c *cityService) DeleteCity(ctx context.Context, id string) error {
	err := c.cityRepository.DeleteCityByID(ctx, id)

	return err
}

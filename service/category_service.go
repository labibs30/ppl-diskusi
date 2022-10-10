package service

import (
	"context"
	"dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"

	"github.com/mashingan/smapping"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category dto.AddCategoryDTO) (entity.Category, error)
	GetAllCategory(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, id string) (entity.Category, error)
	FindByNameOrDesc(ctx context.Context, param string) ([]entity.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	UpdateCategory(ctx context.Context, id string, category dto.UpdateCategoryDTO) (entity.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: repo,
	}
}

func (c *categoryService) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	result, err := c.categoryRepository.GetAllCategory(ctx)

	return result, err
}

func (c *categoryService) CreateCategory(ctx context.Context, category dto.AddCategoryDTO) (entity.Category, error) {
	var createdCategory entity.Category

	err := smapping.FillStruct(&createdCategory, smapping.MapFields(&category))
	if err != nil {
		return entity.Category{}, err
	}

	result, err := c.categoryRepository.InsertCategory(ctx, createdCategory)

	return result, err
}

func (c *categoryService) GetCategoryByID(ctx context.Context, id string) (entity.Category, error) {
	result, err := c.categoryRepository.GetCategoryByID(ctx, id)

	return result, err
}

func (c *categoryService) FindByNameOrDesc(ctx context.Context, param string) ([]entity.Category, error) {
	result, err := c.categoryRepository.GetCategoryByNameOrDesc(ctx, param)

	return result, err
}

func (c *categoryService) UpdateCategory(ctx context.Context, id string, category dto.UpdateCategoryDTO) (entity.Category, error) {
	var updatedCategory entity.Category

	err := smapping.FillStruct(&updatedCategory, smapping.MapFields(&category))
	if err != nil {
		return entity.Category{}, err
	}

	result, err := c.categoryRepository.UpdateCategoryByID(ctx, id, updatedCategory)

	return result, err
}

func (c *categoryService) DeleteCategory(ctx context.Context, id string) error {
	err := c.categoryRepository.DeleteCategoryByID(ctx, id)

	return err
}

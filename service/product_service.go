package service

import (
	"context"
	dto "dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product dto.AddProductDTO) (entity.Product, error)
	GetAllProduct(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id string) (entity.Product, error)
	GetProductByNameOrDesc(ctx context.Context, param string) ([]entity.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	UpdateProduct(ctx context.Context, id string, product dto.UpdateProductDTO) (entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: repo,
	}
}

func (c *productService) CreateProduct(ctx context.Context, product dto.AddProductDTO) (entity.Product, error) {
	var createdProduct entity.Product

	err := smapping.FillStruct(&createdProduct, smapping.MapFields(&product))
	if err != nil {
		return entity.Product{}, err
	}

	result, err := c.productRepository.InsertProduct(ctx, createdProduct)

	return result, err
}

func (c *productService) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	result, err := c.productRepository.GetAllProducts(ctx)

	return result, err
}

func (c *productService) GetProductByID(ctx context.Context, id string) (entity.Product, error) {
	result, err := c.productRepository.GetProductByID(ctx, id)

	return result, err
}

func (c *productService) GetProductByNameOrDesc(ctx context.Context, param string) ([]entity.Product, error) {
	result, err := c.productRepository.GetProductByNameOrDesc(ctx, param)

	return result, err
}

func (c *productService) UpdateProduct(ctx context.Context, id string, product dto.UpdateProductDTO) (entity.Product, error) {
	var updatedProduct entity.Product

	err := smapping.FillStruct(&updatedProduct, smapping.MapFields(&product))
	if err != nil {
		return entity.Product{}, err
	}

	result, err := c.productRepository.UpdateProductByID(ctx, id, updatedProduct)

	return result, err
}

func (c *productService) DeleteProduct(ctx context.Context, id string) error {
	err := c.productRepository.DeleteProductByID(ctx, id)

	return err
}

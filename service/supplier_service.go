package service

import (
	"context"
	dto "dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"

	"github.com/mashingan/smapping"
)

type SupplierService interface {
	CreateSupplier(ctx context.Context, supplier dto.AddSupplierDTO) (entity.Supplier, error)
	GetAllSupplier(ctx context.Context) ([]entity.Supplier, error)
	GetSupplierByID(ctx context.Context, id string) (entity.Supplier, error)
	DeleteSupplier(ctx context.Context, id string) error
	UpdateSupplier(ctx context.Context, id string, supplier dto.UpdateSupplierDTO) (entity.Supplier, error)
}

type supplierService struct {
	supplierRepository repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{
		supplierRepository: repo,
	}
}

func (c *supplierService) CreateSupplier(ctx context.Context, supplier dto.AddSupplierDTO) (entity.Supplier, error) {
	var createdSupplier entity.Supplier

	err := smapping.FillStruct(&createdSupplier, smapping.MapFields(&supplier))
	if err != nil {
		return entity.Supplier{}, err
	}

	result, err := c.supplierRepository.InsertSupplier(ctx, createdSupplier)

	return result, err
}

func (c *supplierService) GetAllSupplier(ctx context.Context) ([]entity.Supplier, error) {
	result, err := c.supplierRepository.GetAllSupplier(ctx)

	return result, err
}

func (c *supplierService) GetSupplierByID(ctx context.Context, id string) (entity.Supplier, error) {
	result, err := c.supplierRepository.GetSupplierByID(ctx, id)

	return result, err
}

func (c *supplierService) UpdateSupplier(ctx context.Context, id string, supplier dto.UpdateSupplierDTO) (entity.Supplier, error) {
	var updatedSupplier entity.Supplier

	err := smapping.FillStruct(&updatedSupplier, smapping.MapFields(&supplier))
	if err != nil {
		return entity.Supplier{}, err
	}

	result, err := c.supplierRepository.UpdateSupplierByID(ctx, id, updatedSupplier)

	return result, err
}

func (c *supplierService) DeleteSupplier(ctx context.Context, id string) error {
	err := c.supplierRepository.DeleteSupplierByID(ctx, id)

	return err
}

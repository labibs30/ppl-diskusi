package service

import (
	"context"
	dto "dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"
	"errors"

	"github.com/mashingan/smapping"
)

type PartnershipService interface {
	CreatePartnership(ctx context.Context, partnership dto.AddPartnershipDTO) (entity.Partnership, error)
	GetAllPartnership(ctx context.Context) ([]entity.Partnership, error)
	GetPartnershipByID(ctx context.Context, id string) (entity.Partnership, error)
	DeletePartnership(ctx context.Context, id string) error
	UpdatePartnership(ctx context.Context, id string, partnership dto.UpdatePartnershipDTO) (entity.Partnership, error)
}

type partnershipService struct {
	partnershipRepository repository.PartnershipRepository
}

func NewPartnershipService(repo repository.PartnershipRepository) PartnershipService {
	return &partnershipService{
		partnershipRepository: repo,
	}
}

func (c *partnershipService) CreatePartnership(ctx context.Context, partnership dto.AddPartnershipDTO) (entity.Partnership, error) {
	var createdPartnership entity.Partnership

	err := smapping.FillStruct(&createdPartnership, smapping.MapFields(&partnership))
	if err != nil {
		return entity.Partnership{}, err
	}

	if (createdPartnership.IsDisetujui >= 1) && (createdPartnership.IsDisetujui <= 3) {
		result, err := c.partnershipRepository.InsertPartnership(ctx, createdPartnership)
		return result, err
	}

	return createdPartnership, errors.New("IsDisetujui value must only 1 / 2 /3")
}

func (c *partnershipService) GetAllPartnership(ctx context.Context) ([]entity.Partnership, error) {
	result, err := c.partnershipRepository.GetAllPartnerships(ctx)

	return result, err
}

func (c *partnershipService) GetPartnershipByID(ctx context.Context, id string) (entity.Partnership, error) {
	result, err := c.partnershipRepository.GetPartnershipByID(ctx, id)

	return result, err
}

func (c *partnershipService) UpdatePartnership(ctx context.Context, id string, partnership dto.UpdatePartnershipDTO) (entity.Partnership, error) {
	var updatedPartnership entity.Partnership

	err := smapping.FillStruct(&updatedPartnership, smapping.MapFields(&partnership))
	if err != nil {
		return entity.Partnership{}, err
	}

	result, err := c.partnershipRepository.UpdatePartnershipByID(ctx, id, updatedPartnership)

	return result, err
}

func (c *partnershipService) DeletePartnership(ctx context.Context, id string) error {
	err := c.partnershipRepository.DeletePartnershipByID(ctx, id)

	return err
}

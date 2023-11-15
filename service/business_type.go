package service

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/entity"
	"github.com/afikrim/go-hexa-template/core/module"
	"github.com/afikrim/go-hexa-template/core/repository"
)

var _ module.BusinessTypeModule = (*businessTypeSvc)(nil)

type businessTypeSvc struct {
	repo repository.BusinessTypeRepository
}

func (b *base) GetBusinessTypeModule() module.BusinessTypeModule {
	return &businessTypeSvc{
		repo: b.repo.GetBusinessTypeRepository(),
	}
}

func (s *businessTypeSvc) GetBusinessTypes(ctx context.Context, in *entity.GetBusinessTypesIn) (*entity.GetBusinessTypesOut, error) {
	businessTypes, err := s.repo.GetBusinessTypes(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.GetBusinessTypesOut{
		BusinessTypes: businessTypes,
	}, nil
}

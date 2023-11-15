package repo

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/entity"
	"github.com/afikrim/go-hexa-template/core/repository"
	"github.com/afikrim/go-hexa-template/repo/dto"
	"gorm.io/gorm"
)

var _ repository.BusinessTypeRepository = (*businessTypeRepo)(nil)

type businessTypeRepo struct {
	db *gorm.DB
}

func (r *base) GetBusinessTypeRepository() repository.BusinessTypeRepository {
	return &businessTypeRepo{
		db: r.db,
	}
}

func (r *businessTypeRepo) GetBusinessTypes(ctx context.Context) (entity.BusinessTypes, error) {
	qdb := r.db.WithContext(ctx).Model(&dto.BusinessTypeDto{})

	var businessTypes dto.BusinessTypesDto
	if err := qdb.Find(&businessTypes).Error; err != nil {
		return nil, err
	}

	return businessTypes.ToEntities(), nil
}

package repo

import (
	"github.com/afikrim/go-hexa-template/core/repository"
	"github.com/afikrim/go-hexa-template/repo/dto"
	"gorm.io/gorm"
)

type (
	base struct {
		db *gorm.DB

		businessTypeRepo repository.BusinessTypeRepository
	}
)

func New(db *gorm.DB) repository.BaseRepository {
	return &base{
		db: db,
	}
}

func (b *base) Migrate() error {
	return b.db.AutoMigrate(
		&dto.BusinessTypeDto{},
	)
}

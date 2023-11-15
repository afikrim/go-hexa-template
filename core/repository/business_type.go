package repository

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/entity"
)

type BusinessTypeRepository interface {
	GetBusinessTypes(ctx context.Context) (entity.BusinessTypes, error)
}

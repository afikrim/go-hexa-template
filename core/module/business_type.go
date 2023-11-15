package module

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/entity"
)

type BusinessTypeModule interface {
	GetBusinessTypes(ctx context.Context, in *entity.GetBusinessTypesIn) (*entity.GetBusinessTypesOut, error)
}

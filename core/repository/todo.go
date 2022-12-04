package repository

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/entity"
)

type TodoRepository interface {
	Create(ctx context.Context, dto *entity.CreateTodoDto) (*entity.Todo, error)
	FindAll(ctx context.Context) ([]entity.Todo, error)
	FindOne(ctx context.Context, id uint64) (*entity.Todo, error)
	Update(ctx context.Context, id uint64, dto *entity.UpdateTodoDto) (*entity.Todo, error)
	Remove(ctx context.Context, id uint64) error
}

package repository

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/entity"
)

type TodoRepository interface {
	Create(ctx context.Context, dto *entity.CreateTodoRequest) (*entity.Todo, error)
	FindAll(ctx context.Context) (entity.Todos, error)
	FindOne(ctx context.Context, id uint64) (*entity.Todo, error)
	Update(ctx context.Context, id uint64, dto *entity.UpdateTodoRequest) (*entity.Todo, error)
	Remove(ctx context.Context, id uint64) error
}

package service

import (
	"context"
	"strconv"

	"github.com/afikrim/go-hexa-template/core/entity"
	repositories "github.com/afikrim/go-hexa-template/core/repository"
)

type TodoService interface {
	Create(ctx context.Context, dto *entity.CreateTodoRequest) (*entity.Todo, error)
	FindAll(ctx context.Context) (entity.Todos, error)
	Update(ctx context.Context, id string, dto *entity.UpdateTodoRequest) (*entity.Todo, error)
	Remove(ctx context.Context, id string) error
}

type service struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, dto *entity.CreateTodoRequest) (*entity.Todo, error) {
	todo, err := s.repo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *service) FindAll(ctx context.Context) (entity.Todos, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) Update(ctx context.Context, id string, dto *entity.UpdateTodoRequest) (*entity.Todo, error) {
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	todo, err := s.repo.Update(ctx, parsedId, dto)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *service) Remove(ctx context.Context, id string) error {
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	return s.repo.Remove(ctx, parsedId)
}

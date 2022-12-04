package todo_repository

import (
	"context"
	"errors"

	"github.com/afikrim/go-hexa-template/core/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, dto *entity.CreateTodoDto) (*entity.Todo, error) {
	todoModel := Todo{
		Title: dto.Title,
	}

	if err := r.db.Create(&todoModel).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return todoModel.ToDomain(), nil
}

func (r *repository) FindAll(ctx context.Context) ([]entity.Todo, error) {
	var todosModel []Todo

	if err := r.db.Find(&todosModel).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	var todos []entity.Todo
	for _, todoModel := range todosModel {
		todos = append(todos, *todoModel.ToDomain())
	}

	return todos, nil
}

func (r *repository) FindOne(ctx context.Context, id uint64) (*entity.Todo, error) {
	var todo Todo

	if err := r.db.First(&todo, id).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

func (r *repository) Update(ctx context.Context, id uint64, dto *entity.UpdateTodoDto) (*entity.Todo, error) {
	todo, err := r.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("Todo not found")
	}

	todoModel := Todo{}.FromDomainWithTimestamps(todo)
	if dto.Title != nil {
		todoModel.Title = *dto.Title
	}
	if dto.Completed != nil {
		todoModel.Completed = *dto.Completed
	}

	if err := r.db.Save(&todoModel).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return todoModel.ToDomain(), nil
}

func (r *repository) Remove(ctx context.Context, id uint64) error {
	todo := entity.Todo{
		ID: id,
	}

	if err := r.db.Delete(&todo).WithContext(ctx).Error; err != nil {
		return err
	}

	return nil
}

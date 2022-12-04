package todo_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/afikrim/go-hexa-template/core/entity"
	errorutil "github.com/afikrim/go-hexa-template/pkg/error"
	replacerutil "github.com/afikrim/go-hexa-template/pkg/replacer"
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

func (r *repository) Create(ctx context.Context, req *entity.CreateTodoRequest) (*entity.Todo, error) {
	todoDto := TodoDto{}.New(req.Title)

	if err := r.db.Create(&todoDto).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return todoDto.ToEntity(), nil
}

func (r *repository) FindAll(ctx context.Context) (entity.Todos, error) {
	var todosDtos []*TodoDto

	if err := r.db.Find(&todosDtos).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	todos := entity.Todos{}
	for _, todoDto := range todosDtos {
		if todoDto == nil {
			continue
		}

		todos = append(todos, todoDto.ToEntity())
	}

	return todos, nil
}

func (r *repository) FindOne(ctx context.Context, id uint64) (*entity.Todo, error) {
	var todo TodoDto

	if err := r.db.First(&todo, id).WithContext(ctx).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("todo not found, %w", errorutil.GENERAL_NOT_FOUND)
		}

		return nil, err
	}

	return todo.ToEntity(), nil
}

func (r *repository) Update(ctx context.Context, id uint64, req *entity.UpdateTodoRequest) (*entity.Todo, error) {
	todo, err := r.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	todoDto := TodoDto{}.FromEntityWithTimestamps(todo)
	todoDto.Title = (replacerutil.Replace(todoDto.Title, req.Title)).(string)
	todoDto.Completed = (replacerutil.Replace(todoDto.Completed, req.Completed)).(bool)

	if err := r.db.Save(&todoDto).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return todoDto.ToEntity(), nil
}

func (r *repository) Remove(ctx context.Context, id uint64) error {
	_, err := r.FindOne(ctx, id)
	if err != nil {
		return err
	}

	todoDto := entity.Todo{
		ID: id,
	}

	if err := r.db.Delete(&todoDto).WithContext(ctx).Error; err != nil {
		return err
	}

	return nil
}

package todo_repository

import (
	"time"

	"github.com/afikrim/go-hexa-template/core/entity"
)

type TodoDto struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	Title     string     `gorm:"column:title;not null"`
	Completed bool       `gorm:"column:completed;not null;default:false"`
	CreatedAt *time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (t TodoDto) TableName() string {
	return "todos"
}

func (t *TodoDto) ToEntity() *entity.Todo {
	return &entity.Todo{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (TodoDto) New(title string) *TodoDto {
	return &TodoDto{
		Title: title,
	}
}

func (TodoDto) FromEntity(d *entity.Todo) *TodoDto {
	return &TodoDto{
		ID:        d.ID,
		Title:     d.Title,
		Completed: d.Completed,
	}
}

func (TodoDto) FromEntityWithTimestamps(d *entity.Todo) *TodoDto {
	parsedCreatedAt, err := time.ParseInLocation("2006-01-02 15:04:05", d.CreatedAt, time.Local)
	if err != nil {
		panic(err)
	}

	parsedUpdatedAt, err := time.ParseInLocation("2006-01-02 15:04:05", d.UpdatedAt, time.Local)
	if err != nil {
		panic(err)
	}

	return &TodoDto{
		ID:        d.ID,
		Title:     d.Title,
		Completed: d.Completed,
		CreatedAt: &parsedCreatedAt,
		UpdatedAt: &parsedUpdatedAt,
	}
}

package dto

import (
	"time"

	"github.com/afikrim/go-hexa-template/core/entity"
)

const (
	BusinessTypeTableName = "business_types"
)

type (
	BusinessTypesDto []*BusinessTypeDto
	BusinessTypeDto  struct {
		ID                 uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
		Serial             string `gorm:"column:serial;type:varchar(255);not null;unique"`
		Name               string `gorm:"column:name;type:varchar(255);not null;unique"`
		Description        string `gorm:"column:description;type:text;not null"`
		BusinessTypeSerial string `gorm:"column:business_type_serial;index"`
		CreatedAt          int64  `gorm:"column:created_at;autoCreateTime;not null;index"`
		UpdatedAt          int64  `gorm:"column:updated_at;autoUpdateTime;not null;index"`
		DeletedAt          int64  `gorm:"column:deleted_at;index"`
	}
)

func (BusinessTypeDto) TableName() string {
	return BusinessTypeTableName
}

func (bt *BusinessTypeDto) FromEntity(e *entity.BusinessType) {
	if e == nil {
		return
	}

	bt.ID = e.ID
	bt.Serial = e.Serial
	bt.Name = e.Name
	bt.Description = e.Description
	bt.BusinessTypeSerial = e.BusinessTypeSerial
	bt.CreatedAt = e.CreatedAt.Unix()
	bt.UpdatedAt = e.UpdatedAt.Unix()
	if e.DeletedAt != nil {
		bt.DeletedAt = e.DeletedAt.Unix()
	}
}

func (bt *BusinessTypeDto) ToEntity() *entity.BusinessType {
	if bt == nil {
		return nil
	}

	deletedAt := time.Unix(bt.DeletedAt, 0)
	return &entity.BusinessType{
		ID:                 bt.ID,
		Serial:             bt.Serial,
		Name:               bt.Name,
		Description:        bt.Description,
		BusinessTypeSerial: bt.BusinessTypeSerial,
		CreatedAt:          time.Unix(bt.CreatedAt, 0),
		UpdatedAt:          time.Unix(bt.UpdatedAt, 0),
		DeletedAt:          &deletedAt,
	}
}

func (bts BusinessTypesDto) ToEntities() entity.BusinessTypes {
	if len(bts) == 0 {
		return nil
	}

	var entities entity.BusinessTypes
	for _, bt := range bts {
		entities = append(entities, bt.ToEntity())
	}

	return entities
}

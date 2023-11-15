package entity

import "time"

type (
	BusinessTypes []*BusinessType
	BusinessType  struct {
		ID                 uint64     `json:"id"`
		Serial             string     `json:"serial"`
		Name               string     `json:"name"`
		Description        string     `json:"description"`
		BusinessTypeSerial string     `json:"business_type_serial"`
		CreatedAt          time.Time  `json:"created_at"`
		UpdatedAt          time.Time  `json:"updated_at"`
		DeletedAt          *time.Time `json:"deleted_at"`
	}

	GetBusinessTypesIn struct{}

	GetBusinessTypesOut struct {
		BusinessTypes BusinessTypes `json:"business_types"`
	}
)

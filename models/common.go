package models

import (
	"time"
)

// MyGormModel mimixks GormModel but uses uuid's for ID, generated in go
type MyGormModel struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TimestampModel ...
type TimestampModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

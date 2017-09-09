package models

import (
	"time"
)

type Default struct {
	ID        uint64      `gorm:"primary_key"`
	CreatedAt time.Time   `gorm:"type:timestamp(3);not null;DEFAULT:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time   `gorm:"type:timestamp(3);not null;DEFAULT:CURRENT_TIMESTAMP(3)"`
}

type Memo struct {
	Default
	Text string `gorm:"type:text"`
}

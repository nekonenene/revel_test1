package models

import (
	"github.com/jinzhu/gorm"
)

type Memo struct {
	gorm.Model
	Text string
}

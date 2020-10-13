package models

import (
	"github.com/jinzhu/gorm"
)

type Positions struct {
	gorm.Model
	UserID  string
	PosName string
}

package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID      string
	PrenameTh   string
	FirstnameTh string
	LastnameTh  string
	PrenameEn   string
	FirstnameEn string
	LastnameEn  string
	Password    string
	Gender      int
	Department  int
	Status      int
	Position    int
	Img         string
}

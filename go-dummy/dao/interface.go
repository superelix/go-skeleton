package dao

import "github.com/jinzhu/gorm"

type CreateInterface interface {
	Create(db *gorm.DB) (err error)
}

func Create(ci CreateInterface, db *gorm.DB) (err error) {
	return ci.Create(db)
}

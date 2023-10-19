package dao

import (
	"go-dummy-project/go-dummy/common"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

func (ab *DummyUser) BeforeCreate(scope *gorm.Scope) {
	pk := DummyUserPrefix + common.CreateUUID()
	scope.SetColumn("ID", pk)
}

func (u *DummyUser) Create(db *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	err = db.Create(u).Error
	return err
}

package dao

import "go-dummy-project/go-dummy/common"

const DummyUserPrefix = "dup_"

type DummyUser struct {
	common.BaseModel
	ID        string `gorm:"primaryKey;column:id" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"coulmn:last_name" json:"last_name"`
	Email     string `gorm:"column:email" json:"email,omitempty"`
	Code      int32  `gorm:"column:code" json:"code"`
}

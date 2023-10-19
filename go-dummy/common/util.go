package common

import (
	"path/filepath"
	"runtime"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

type PropertyMap map[string]interface{}

func ProjectRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	rootPath := strings.Split(filepath.Dir(b), "/common")
	return rootPath[0] + "/"
}

func CreateUUID() string {
	return uuid.NewV4().String()
}

package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Base entity
type Base struct {
	ID        string    `json:"id" gorm:"column:id;type:uuid;not null" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null" valid:"notnull"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null" valid:"-"`
}

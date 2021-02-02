package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// User entity
type User struct {
	Base  `valid:"required"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {

	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}

	return nil
}

// NewUser create and return a new User
func NewUser(name string, email string) (*User, error) {

	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	if err := user.isValid(); err != nil {
		return nil, err
	}

	return &user, nil
}

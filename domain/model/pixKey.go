package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// IPixKeyRepository is a interface that shoud implement by the repository layer
type IPixKeyRepository interface {
	RegisterPixKey(pixKey *PixKey) (*PixKey, error)
	FindPixKeyByID(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	FindBankByID(id string) (*Bank, error)
	AddAccount(account *Account) error
	FindAccountByID(id string) (*Account, error)
}

// PixKey Entity
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:varchar(20);not null" valid:"notnull"`
	Key       string   `json:"key" gorm:"type:varchar(20);not null" valid:"notnull"`
	AccountID string   `json:"account_id" gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"_"`
	Status    string   `json:"status" gorm:"type:varchar(20);not null" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {

	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("Invalid type of Key")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("Invalid Status")
	}

	if err != nil {
		return err
	}

	return nil
}

// NewPixKey create and return a new PixKey
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {

	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	if err := pixKey.isValid(); err != nil {
		return nil, err
	}

	return &pixKey, nil
}

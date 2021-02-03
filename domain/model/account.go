package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// Account Entity
type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	BankID    string    `json:"bank_id" gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKey    []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func (acc *Account) isValid() error {

	if _, err := govalidator.ValidateStruct(acc); err != nil {
		return err
	}

	return nil
}

// NewAccount create and return a new Account
func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {

	account := Account{
		Bank:      bank,
		Number:    number,
		OwnerName: ownerName,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return &account, nil
}

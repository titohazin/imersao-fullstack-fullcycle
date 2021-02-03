package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/titohazin/imersao-fullstack-fullcycle/domain/model"
)

// PixKeyRepositoryDb is the ORM implementation of the Pix Key
type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

// AddBank is a interface implementation of the AddBank in Pix Key
func (repo *PixKeyRepositoryDb) AddBank(bank *model.Bank) error {

	if err := repo.Db.Create(bank).Error; err != nil {
		return err
	}

	return nil
}

// AddAccount is a interface implementation of the AddAccount in Pix Key
func (repo *PixKeyRepositoryDb) AddAccount(account *model.Account) error {

	if err := repo.Db.Create(account).Error; err != nil {
		return err
	}

	return nil
}

// RegisterPixKey is a interface implementation of the RegisterPixKey in Pix Key
func (repo *PixKeyRepositoryDb) RegisterPixKey(pixKey *model.PixKey) (*model.PixKey, error) {

	if err := repo.Db.Create(pixKey).Error; err != nil {
		return nil, err
	}

	return pixKey, nil
}

// FindPixKeyByID is a interface implementation of the FindPixKeyById in Pix Key
func (repo *PixKeyRepositoryDb) FindPixKeyByID(pixKey string, kind string) (*model.PixKey, error) {

	var pixKeyResult model.PixKey

	repo.Db.Preload("Account.Bank").First(&pixKeyResult, "kind = ? and key = ?", kind, pixKey)

	if pixKeyResult.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKeyResult, nil
}

// FindAccountByID is a interface implementation of the FindAccountByID in Pix Key
func (repo *PixKeyRepositoryDb) FindAccountByID(id string) (*model.Account, error) {

	var accountResult model.Account

	repo.Db.Preload("Bank").First(&accountResult, "id = ?", id)

	if accountResult.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}

	return &accountResult, nil
}

// FindBankByID is a interface implementation of the FindBankByID in Pix Key
func (repo *PixKeyRepositoryDb) FindBankByID(id string) (*model.Bank, error) {

	var bankResult model.Bank

	repo.Db.First(&bankResult, "id = ?", id)

	if bankResult.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}

	return &bankResult, nil
}

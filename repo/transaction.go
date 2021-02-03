package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/titohazin/imersao-fullstack-fullcycle/domain/model"
)

// TransactionRepositoryDb is the ORM implementation of the Transaction
type TransactionRepositoryDb struct {
	Db *gorm.DB
}

// Register is a interface implementation of the Register in Transaction
func (repo *TransactionRepositoryDb) Register(transaction *model.Transaction) (*model.Transaction, error) {

	if err := repo.Db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

// Save is a interface implementation of the Save in Transaction
func (repo *TransactionRepositoryDb) Save(transaction *model.Transaction) (*model.Transaction, error) {

	if err := repo.Db.Save(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

// Find is a interface implementation of the Find in Transaction
func (repo *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {

	var transactionResult model.Transaction

	repo.Db.Preload("AccountFrom.Bank").First(&transactionResult, "id = ?", id)

	if transactionResult.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transactionResult, nil
}

package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// Transaction Constants Status
const (
	TransactionPeding    string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

// ITransactionData is a interface that shoud implement by the repository layer
type ITransactionData interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

// Transactions is Transaction list
type Transactions struct {
	Transactions []*Transaction
}

// Transaction of the application
type Transaction struct {
	Base              `valid:"required"`
	AccountIDFrom     string   `json:"account_id_from" gorm:"column:account_id_from;type:uuid;not null" valid:"-"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" gorm:"type:float;not null" valid:"notnull"`
	PixKeyIDTo        string   `json:"pix_key_id_to" gorm:"column:pix_key_id_to;type:uuid;not null" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" gorm:"type:varchar(20);not null" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255);not null" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255" valid:"-"`
}

func (transaction *Transaction) isValid() error {

	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("Amount must be greater than 0")
	}

	if transaction.Status != TransactionPeding &&
		transaction.Status != TransactionCompleted &&
		transaction.Status != TransactionError {
		return errors.New("Invalid Status for the transaction")
	}

	if transaction.PixKeyTo.Account.ID == transaction.AccountFrom.ID {
		return errors.New("The source and destination accounts can not be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

// NewTransaction create and return a new Transaction
func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {

	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPeding,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	if err := transaction.isValid(); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Complete - change Transaction Status to completed
func (transaction *Transaction) Complete() error {
	return changeTransactionStatus(transaction, TransactionCompleted)
}

// Confirm - change Transaction Status to confimed
func (transaction *Transaction) Confirm() error {
	return changeTransactionStatus(transaction, TransactionConfirmed)
}

// Cancel - change Transaction Status to error
func (transaction *Transaction) Cancel(cancelDescription string) error {
	transaction.CancelDescription = cancelDescription
	return changeTransactionStatus(transaction, TransactionError)
}

func changeTransactionStatus(transaction *Transaction, status string) error {
	transaction.Status = status
	transaction.UpdatedAt = time.Now()
	err := transaction.isValid()
	return err
}

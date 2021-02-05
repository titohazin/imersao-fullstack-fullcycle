package model

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

// Transaction model
type Transaction struct {
	ID           string  `json:"id" validate:"required, uuid4"`
	AccountID    string  `json:"account_id" validate:"required, uuid4"`
	Amount       float64 `json:"amount" validate:"required, numeric"`
	PixKeyTo     string  `json:"pix_key_to" validate:"required"`
	PixKeyToKind string  `json:"pix_key_to_kind" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error" validate:"-"`
}

func (transaction *Transaction) isValid() error {

	if _, err := govalidator.ValidateStruct(transaction); err != nil {
		return err
	}

	return nil
}

// JSONToModel Parse data JSON to Transaction Model
func (transaction *Transaction) JSONToModel(data []byte) error {

	err := json.Unmarshal(data, transaction)

	if err != nil {
		return err
	}

	err = transaction.isValid()

	if err != nil {
		return err
	}

	return nil
}

// ModelToJSON Parse Transaction Model to JSON
func (transaction *Transaction) ModelToJSON() ([]byte, error) {

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(transaction)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewTransaction return a empty Transaction
func NewTransaction() *Transaction {
	return &Transaction{}
}

package usecase

import (
	"errors"

	"github.com/titohazin/imersao-fullstack-fullcycle/domain/model"
)

// TransactionUseCase Transaction use case
type TransactionUseCase struct {
	TransactionRepository model.ITransactionRepository
	PixKeyRepository      model.IPixKeyRepository
}

// Register a new transaction
func (transUseCase *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyToKind string, description string) (*model.Transaction, error) {

	account, err := transUseCase.PixKeyRepository.FindAccountByID(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := transUseCase.PixKeyRepository.FindPixKeyByID(pixKeyTo, pixKeyToKind)

	if err != nil {
		return nil, err
	}

	newTransaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	err = transUseCase.TransactionRepository.Save(newTransaction)

	if err != nil || newTransaction.ID == "" {
		return nil, errors.New("unable to process this Transaction")
	}

	return newTransaction, nil
}

// Confirm a transaction
func (transUseCase *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {

	transaction, err := transUseCase.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	err = transaction.Confirm()

	if err != nil {
		return nil, err
	}

	err = transUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Complete a transaction
func (transUseCase *TransactionUseCase) Complete(transactionID string) (*model.Transaction, error) {

	transaction, err := transUseCase.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	err = transaction.Complete()

	if err != nil {
		return nil, err
	}

	err = transUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Cancel a transaction
func (transUseCase *TransactionUseCase) Cancel(transactionID string, reason string) (*model.Transaction, error) {

	transaction, err := transUseCase.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	err = transaction.Cancel(reason)

	if err != nil {
		return nil, err
	}

	err = transUseCase.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

package usecase

import "github.com/titohazin/imersao-fullstack-fullcycle/domain/model"

// PixUseCase Pix use case
type PixUseCase struct {
	PixKeyRepository model.IPixKeyData
}

// RegisterPixKey Register PixKey
func (pixUseCase *PixUseCase) RegisterPixKey(key string, kind string, accountID string) (*model.PixKey, error) {

	account, err := pixUseCase.PixKeyRepository.FindAccountByID(accountID)

	if err != nil {
		return nil, err
	}

	newPixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	_, err = pixUseCase.PixKeyRepository.RegisterPixKey(newPixKey)

	if err != nil {
		return nil, err
	}

	return newPixKey, nil
}

// FindPixKey Find PixKey
func (pixUseCase *PixUseCase) FindPixKey(key string, kind string) (*model.PixKey, error) {

	pixkey, err := pixUseCase.FindPixKey(key, kind)

	if err != nil {
		return nil, err
	}

	return pixkey, nil
}

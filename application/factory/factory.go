package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/usecase"
	"github.com/titohazin/imersao-fullstack-fullcycle/repo"
)

// TranactionUseCaseFactory Tranaction Use Case Factory
func TranactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {

	pixKeyRepository := repo.PixKeyRepositoryDb{Db: database}
	transactionRepository := repo.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		PixKeyRepository:      &pixKeyRepository,
		TransactionRepository: &transactionRepository,
	}

	return transactionUseCase
}

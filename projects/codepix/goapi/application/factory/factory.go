package factory

import "gituhb.com/jinzhu/gorm"

func TransactionUseCaseFctory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepo := PixKeyRepositoryDb{Db: database}
	transactionRepo := TransactionRepositoryDb{Db: database}
	
	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepo,
		PixRepository: &pixRepo
	}

	return tranactionUseCase
}
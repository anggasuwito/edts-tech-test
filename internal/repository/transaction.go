package repository

import (
	"context"
	"edts-tech-test/internal/domain/model"
	"edts-tech-test/internal/utils"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, data *model.Transaction) error
}

type transactionRepo struct {
	masterDB *gorm.DB
}

func NewTransactionRepo(masterDB *gorm.DB) TransactionRepo {
	return &transactionRepo{
		masterDB: masterDB,
	}
}

func (r *transactionRepo) useTX(ctx context.Context) *gorm.DB {
	if tx := utils.GetTransactionFromContext(ctx); tx != nil {
		return tx
	}
	return r.masterDB
}

func (r *transactionRepo) CreateTransaction(ctx context.Context, data *model.Transaction) error {
	db := r.useTX(ctx)
	err := db.Debug().Create(data).Error
	if err != nil {
		return utils.ErrInternal("Failed create new transaction : "+err.Error(), "transactionRepo.CreateTransaction")
	}
	return nil
}

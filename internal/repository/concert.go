package repository

import (
	"context"
	"edts-tech-test/internal/utils"
	"gorm.io/gorm"
)

type ConcertRepo interface {
}

type concertRepo struct {
	masterDB *gorm.DB
}

func NewConcertRepo(masterDB *gorm.DB) ConcertRepo {
	return &concertRepo{
		masterDB: masterDB,
	}
}

func (r *concertRepo) useTX(ctx context.Context) *gorm.DB {
	if tx := utils.GetTransactionFromContext(ctx); tx != nil {
		return tx
	}
	return r.masterDB
}

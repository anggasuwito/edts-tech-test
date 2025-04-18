package usecase

import (
	"context"
	"edts-tech-test/internal/domain/entity"
	"edts-tech-test/internal/repository"
)

type ConcertUC interface {
	GetConcertList(ctx context.Context, req *entity.GetConcertListRequest) (*entity.GetConcertListResponse, error)
	BookingConcert(ctx context.Context, req *entity.BookingConcertRequest) (*entity.BookingConcertResponse, error)
	GetConcertPurchaseHistory(ctx context.Context, req *entity.GetConcertPurchaseHistoryRequest) (*entity.GetConcertPurchaseHistoryResponse, error)
}

type concertUC struct {
	txWrapper   repository.TransactionWrapper
	concertRepo repository.ConcertRepo
}

func NewConcertUC(
	txWrapper repository.TransactionWrapper,
	concertRepo repository.ConcertRepo,
) ConcertUC {
	return &concertUC{
		txWrapper:   txWrapper,
		concertRepo: concertRepo,
	}
}

func (u *concertUC) GetConcertList(ctx context.Context, req *entity.GetConcertListRequest) (*entity.GetConcertListResponse, error) {
	return &entity.GetConcertListResponse{}, nil
}

func (u *concertUC) BookingConcert(ctx context.Context, req *entity.BookingConcertRequest) (*entity.BookingConcertResponse, error) {
	return &entity.BookingConcertResponse{}, nil
}

func (u *concertUC) GetConcertPurchaseHistory(ctx context.Context, req *entity.GetConcertPurchaseHistoryRequest) (*entity.GetConcertPurchaseHistoryResponse, error) {
	return &entity.GetConcertPurchaseHistoryResponse{}, nil
}

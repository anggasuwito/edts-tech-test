package usecase

import (
	"context"
	"edts-tech-test/internal/domain/entity"
	"edts-tech-test/internal/repository"
)

type UserUC interface {
	GetUserPurchaseHistory(ctx context.Context, req *entity.GetUserPurchaseHistoryRequest) (*entity.GetUserPurchaseHistoryResponse, error)
}

type userUC struct {
	concertRepo repository.ConcertRepo
}

func NewUserUC(
	concertRepo repository.ConcertRepo,
) UserUC {
	return &userUC{
		concertRepo: concertRepo,
	}
}

func (u *userUC) GetUserPurchaseHistory(ctx context.Context, req *entity.GetUserPurchaseHistoryRequest) (*entity.GetUserPurchaseHistoryResponse, error) {
	return &entity.GetUserPurchaseHistoryResponse{}, nil
}

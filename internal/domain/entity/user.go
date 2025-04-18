package entity

type (
	GetUserPurchaseHistoryRequest struct {
		*ListPaginationRequest
	}

	GetUserPurchaseHistoryResponse struct {
		*ListPaginationResponse
		Data []*ConcertPurchaseHistory `json:"data"`
	}
)

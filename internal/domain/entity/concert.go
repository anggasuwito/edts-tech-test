package entity

type (
	GetConcertListRequest struct {
		*ListPaginationRequest
	}

	GetConcertListResponse struct {
		*ListPaginationResponse
		Data []*Concert `json:"data"`
	}

	BookingConcertRequest struct {
		UserPhone string `json:"user_phone"`
		ConcertID string `json:"concert_id"`
		Qty       int64  `json:"qty"`
	}

	BookingConcertResponse struct {
		*ConcertPurchaseHistory
	}

	GetConcertPurchaseHistoryRequest struct {
		ConcertID string `json:"concert_id"`
	}

	GetConcertPurchaseHistoryResponse struct {
		*ListPaginationResponse
		Data []*ConcertPurchaseHistory `json:"data"`
	}

	Concert struct {
		ID              string `json:"id"`
		AvailableFrom   string `json:"available_from"`
		AvailableTo     string `json:"available_to"`
		AvailableTicket int64  `json:"available_ticket"`
		PlayAt          string `json:"play_at"`
		Name            string `json:"name"`
		Price           int64  `json:"price"`
		CreatedAt       string `json:"created_at"`
	}

	ConcertPurchaseHistory struct {
		ID          string `json:"id"`
		ConcertID   string `json:"concert_id"`
		UserPhone   string `json:"user_phone"`
		ConcertName string `json:"concert_name"`
		Price       int64  `json:"price"`
		Qty         int64  `json:"qty"`
		TotalPrice  int64  `json:"total_price"`
	}
)

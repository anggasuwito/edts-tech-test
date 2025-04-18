package model

func (m *ConcertPurchaseHistory) TableName() string {
	return "concert_purchase_history"
}

type ConcertPurchaseHistory struct {
	BaseModel

	ConcertID   string `gorm:"column:concert_id;size:36;"`
	UserPhone   string `gorm:"column:user_phone;size:100;"`
	ConcertName string `gorm:"column:concert_name;size:255;"`
	Price       int64  `gorm:"column:price"`
	Qty         int64  `gorm:"column:qty"`
	TotalPrice  int64  `gorm:"column:total_price"`
}

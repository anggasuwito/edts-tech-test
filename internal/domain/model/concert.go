package model

import (
	"database/sql"
)

func (m *Concert) TableName() string {
	return "concert"
}

type Concert struct {
	BaseModel

	AvailableFrom   sql.NullTime `gorm:"column:available_from"`
	AvailableTo     sql.NullTime `gorm:"column:available_to"`
	AvailableTicket int64        `gorm:"column:available_ticket"`
	PlayAt          sql.NullTime `gorm:"column:play_at"`
	Price           int64        `gorm:"column:price"`
	Name            string       `gorm:"column:name;size:255;"`
}

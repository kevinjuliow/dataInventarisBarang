package domain

import (
	"database/sql"
	"time"
)

type ActivityLog struct {
	Id           int            `json:"id"`
	ItemId       int            `json:"item_id"`
	UserId       int            `json:"user_id"`
	Tipe         string         `json:"tipe"`
	Jumlah       int            `json:"jumlah"`
	StockSebelum int            `json:"stock_sebelum"`
	StockSesudah int            `json:"stock_sesudah"`
	Keterangan   sql.NullString `json:"keterangan"`
	CreatedAt    time.Time      `json:"created_at"`
}

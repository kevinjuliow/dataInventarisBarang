package domain

import (
	"database/sql"
	"time"
)

type Item struct {
	Id         int            `json:"id"`
	Nama       string         `json:"nama"`
	Keterangan sql.NullString `json:"keterangan"`
	Stock      int            `json:"stock"`
	Harga      float64        `json:"harga"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

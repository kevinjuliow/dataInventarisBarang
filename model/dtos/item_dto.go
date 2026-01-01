package dtos

import "time"

type ItemCreateRequest struct {
	Nama       string  `validate:"required,min=1,max=200" json:"nama"`
	Keterangan string  `json:"keterangan"`
	Stock      int     `validate:"required,min=0" json:"stock"`
	Harga      float64 `validate:"required,min=0" json:"harga"`
}

type ItemUpdateRequest struct {
	Id         int     `validate:"required" json:"id"`
	Nama       string  `validate:"required,min=1,max=200" json:"nama"`
	Keterangan string  `json:"keterangan"`
	Stock      int     `validate:"required,min=0" json:"stock"`
	Harga      float64 `validate:"required,min=0" json:"harga"`
}

type ItemResponse struct {
	Id         int       `json:"id"`
	Nama       string    `json:"nama"`
	Keterangan string    `json:"keterangan"`
	Stock      int       `json:"stock"`
	Harga      float64   `json:"harga"`
	CreatedAt  time.Time `json:"created_at"`
}

package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
)

type ItemRepositoryImpl struct {
}

func NewItemRepository() ItemRepository {
	return &ItemRepositoryImpl{}
}

func (repository *ItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item {
	SQL := "INSERT INTO items(nama, keterangan, stock, harga) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, item.Nama, item.Keterangan, item.Stock, item.Harga)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	item.Id = int(id)
	return item
}

func (repository *ItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item {
	SQL := "UPDATE items SET nama = ?, keterangan = ?, stock = ?, harga = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, item.Nama, item.Keterangan, item.Stock, item.Harga, item.Id)
	helper.PanicIfError(err)

	return item
}

func (repository *ItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, itemId int) {
	SQL := "DELETE FROM items WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, itemId)
	helper.PanicIfError(err)
}

func (repository *ItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, itemId int) (domain.Item, error) {
	SQL := "SELECT id, nama, keterangan, stock, harga, created_at, updated_at FROM items WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, itemId)
	helper.PanicIfError(err)
	defer rows.Close()

	item := domain.Item{}
	if rows.Next() {
		err := rows.Scan(&item.Id, &item.Nama, &item.Keterangan, &item.Stock, &item.Harga, &item.CreatedAt, &item.UpdatedAt)
		helper.PanicIfError(err)
		return item, nil
	} else {
		return item, errors.New("item not found")
	}
}

func (repository *ItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Item {
	SQL := "SELECT id, nama, keterangan, stock, harga, created_at, updated_at FROM items"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		item := domain.Item{}
		err := rows.Scan(&item.Id, &item.Nama, &item.Keterangan, &item.Stock, &item.Harga, &item.CreatedAt, &item.UpdatedAt)
		helper.PanicIfError(err)
		items = append(items, item)
	}
	return items
}

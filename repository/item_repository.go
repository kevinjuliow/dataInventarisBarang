package repository

import (
	"context"
	"database/sql"

	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
)

type ItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item
	Update(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item
	Delete(ctx context.Context, tx *sql.Tx, itemId int)
	FindById(ctx context.Context, tx *sql.Tx, itemId int) (domain.Item, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Item
}

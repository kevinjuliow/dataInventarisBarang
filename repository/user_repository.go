package repository

import (
	"context"
	"database/sql"

	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
}

package repository

import (
	"context"
	"database/sql"

	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
)

type LogRepository interface {
	Save(ctx context.Context, tx *sql.Tx, log domain.ActivityLog) domain.ActivityLog
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ActivityLog
}

type LogRepositoryImpl struct {
}

func NewLogRepository() LogRepository {
	return &LogRepositoryImpl{}
}

func (repository *LogRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, log domain.ActivityLog) domain.ActivityLog {
	SQL := `INSERT INTO activity_log(item_id, user_id, tipe, jumlah, stock_sebelum, stock_sesudah, keterangan) 
			VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := tx.ExecContext(ctx, SQL,
		log.ItemId,
		log.UserId,
		log.Tipe,
		log.Jumlah,
		log.StockSebelum,
		log.StockSesudah,
		log.Keterangan,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	log.Id = int(id)
	return log
}

func (repository *LogRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ActivityLog {
	SQL := "SELECT id, item_id, user_id, tipe, jumlah, stock_sebelum, stock_sesudah, keterangan, created_at FROM activity_log"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var logs []domain.ActivityLog
	for rows.Next() {
		log := domain.ActivityLog{}
		err := rows.Scan(
			&log.Id,
			&log.ItemId,
			&log.UserId,
			&log.Tipe,
			&log.Jumlah,
			&log.StockSebelum,
			&log.StockSesudah,
			&log.Keterangan,
			&log.CreatedAt,
		)
		helper.PanicIfError(err)
		logs = append(logs, log)
	}
	return logs
}

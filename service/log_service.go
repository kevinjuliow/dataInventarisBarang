package service

import (
	"context"
	"database/sql"

	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
	"github.com/kevinjuliow/dataInventarisBarang/repository"
)

type LogService interface {
	FindAll(ctx context.Context) []domain.ActivityLog
}

type LogServiceImpl struct {
	LogRepository repository.LogRepository
	DB            *sql.DB
}

func NewLogService(logRepository repository.LogRepository, db *sql.DB) LogService {
	return &LogServiceImpl{
		LogRepository: logRepository,
		DB:            db,
	}
}

func (service *LogServiceImpl) FindAll(ctx context.Context) []domain.ActivityLog {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	return service.LogRepository.FindAll(ctx, tx)
}

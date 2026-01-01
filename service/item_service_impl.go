package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/kevinjuliow/dataInventarisBarang/exception"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
	"github.com/kevinjuliow/dataInventarisBarang/repository"
)

type ItemServiceImpl struct {
	ItemRepository repository.ItemRepository
	LogRepository  repository.LogRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewItemService(itemRepository repository.ItemRepository, logRepository repository.LogRepository, db *sql.DB, validate *validator.Validate) ItemService {
	return &ItemServiceImpl{
		ItemRepository: itemRepository,
		LogRepository:  logRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *ItemServiceImpl) Create(ctx context.Context, request dtos.ItemCreateRequest, userId int) dtos.ItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	item := domain.Item{
		Nama:  request.Nama,
		Stock: request.Stock,
		Harga: request.Harga,
		Keterangan: sql.NullString{
			String: request.Keterangan,
			Valid:  request.Keterangan != "",
		},
	}

	item = service.ItemRepository.Save(ctx, tx, item)

	if item.Stock > 0 {
		log := domain.ActivityLog{
			ItemId:       item.Id,
			UserId:       userId,
			Tipe:         "IN",
			Jumlah:       item.Stock,
			StockSebelum: 0,
			StockSesudah: item.Stock,
			Keterangan:   sql.NullString{String: "Stock awal", Valid: true},
		}
		service.LogRepository.Save(ctx, tx, log)
	}

	return dtos.ItemResponse{
		Id:         item.Id,
		Nama:       item.Nama,
		Keterangan: item.Keterangan.String,
		Stock:      item.Stock,
		Harga:      item.Harga,
		CreatedAt:  item.CreatedAt,
	}
}

func (service *ItemServiceImpl) Update(ctx context.Context, request dtos.ItemUpdateRequest, userId int) dtos.ItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	item, err := service.ItemRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	oldStock := item.Stock

	item.Nama = request.Nama
	item.Stock = request.Stock
	item.Harga = request.Harga
	item.Keterangan = sql.NullString{
		String: request.Keterangan,
		Valid:  request.Keterangan != "",
	}

	item = service.ItemRepository.Update(ctx, tx, item)

	if item.Stock != oldStock {
		tipe := "IN"
		jumlah := item.Stock - oldStock

		if item.Stock < oldStock {
			tipe = "OUT"
			jumlah = oldStock - item.Stock
		}

		log := domain.ActivityLog{
			ItemId:       item.Id,
			UserId:       userId,
			Tipe:         tipe,
			Jumlah:       jumlah,
			StockSebelum: oldStock,
			StockSesudah: item.Stock,
			Keterangan:   sql.NullString{String: "Update Stock", Valid: true},
		}
		service.LogRepository.Save(ctx, tx, log)
	}

	return dtos.ItemResponse{
		Id:         item.Id,
		Nama:       item.Nama,
		Keterangan: item.Keterangan.String,
		Stock:      item.Stock,
		Harga:      item.Harga,
		CreatedAt:  item.CreatedAt,
	}
}

func (service *ItemServiceImpl) Delete(ctx context.Context, itemId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	item, err := service.ItemRepository.FindById(ctx, tx, itemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ItemRepository.Delete(ctx, tx, item.Id)
}

func (service *ItemServiceImpl) FindById(ctx context.Context, itemId int) dtos.ItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	item, err := service.ItemRepository.FindById(ctx, tx, itemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return dtos.ItemResponse{
		Id:         item.Id,
		Nama:       item.Nama,
		Keterangan: item.Keterangan.String,
		Stock:      item.Stock,
		Harga:      item.Harga,
		CreatedAt:  item.CreatedAt,
	}
}

func (service *ItemServiceImpl) FindAll(ctx context.Context) []dtos.ItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	items := service.ItemRepository.FindAll(ctx, tx)

	var itemResponses []dtos.ItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, dtos.ItemResponse{
			Id:         item.Id,
			Nama:       item.Nama,
			Keterangan: item.Keterangan.String,
			Stock:      item.Stock,
			Harga:      item.Harga,
			CreatedAt:  item.CreatedAt,
		})
	}
	return itemResponses
}

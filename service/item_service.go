package service

import (
	"context"

	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
)

type ItemService interface {
	Create(ctx context.Context, request dtos.ItemCreateRequest, userId int) dtos.ItemResponse
	Update(ctx context.Context, request dtos.ItemUpdateRequest, userId int) dtos.ItemResponse
	Delete(ctx context.Context, itemId int)
	FindById(ctx context.Context, itemId int) dtos.ItemResponse
	FindAll(ctx context.Context) []dtos.ItemResponse
}

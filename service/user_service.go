package service

import (
	"context"

	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
)

type UserService interface {
	Register(ctx context.Context, request dtos.UserRegisterRequest) dtos.UserResponse
	Login(ctx context.Context, request dtos.UserLoginRequest) string
}

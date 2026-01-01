package service

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/domain"
	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
	"github.com/kevinjuliow/dataInventarisBarang/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request dtos.UserRegisterRequest) dtos.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	passwordHash := string(bytes)

	user := domain.User{
		Email:    request.Email,
		Password: passwordHash,
		Name:     request.Name,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return dtos.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, request dtos.UserLoginRequest) string {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(errors.New("invalid email or password"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(errors.New("invalid email or password"))
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		panic(errors.New("error fetching data in JWT_SECRET"))
	}

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	helper.PanicIfError(err)

	return t
}

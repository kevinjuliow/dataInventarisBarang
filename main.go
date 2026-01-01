package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/kevinjuliow/dataInventarisBarang/app"
	"github.com/kevinjuliow/dataInventarisBarang/controller"
	"github.com/kevinjuliow/dataInventarisBarang/exception"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/middleware"
	"github.com/kevinjuliow/dataInventarisBarang/repository"
	"github.com/kevinjuliow/dataInventarisBarang/service"
)

func main() {
	db := app.NewDb()
	validate := validator.New()

	userRepo := repository.NewUserRepositoryImpl()
	logRepo := repository.NewLogRepository()
	userService := service.NewUserService(userRepo, db, validate)
	userController := controller.NewUserController(userService)

	itemRepo := repository.NewItemRepository()
	itemService := service.NewItemService(itemRepo, logRepo, db, validate)
	itemController := controller.NewItemController(itemService)

	logService := service.NewLogService(logRepo, db)
	logController := controller.NewLogController(logService)

	router := httprouter.New()

	router.POST("/api/register", userController.Register)
	router.POST("/api/login", userController.Login)

	router.GET("/api/items", itemController.FindAll)
	router.GET("/api/items/:itemId", itemController.FindById)
	router.POST("/api/items", itemController.Create)
	router.PUT("/api/items/:itemId", itemController.Update)
	router.DELETE("/api/items/:itemId", itemController.Delete)

	router.GET("/api/logs", logController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

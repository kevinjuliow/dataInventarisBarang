package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
	"github.com/kevinjuliow/dataInventarisBarang/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := dtos.UserRegisterRequest{}
	helper.ReadFromRequestBody(request, &userRegisterRequest)

	userResponse := controller.UserService.Register(request.Context(), userRegisterRequest)

	webResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := dtos.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	token := controller.UserService.Login(request.Context(), userLoginRequest)

	webResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data: map[string]string{
			"token": token,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

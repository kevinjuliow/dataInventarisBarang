package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
	"github.com/kevinjuliow/dataInventarisBarang/service"
)

type ItemControllerImpl struct {
	ItemService service.ItemService
}

func NewItemController(itemService service.ItemService) ItemController {
	return &ItemControllerImpl{
		ItemService: itemService,
	}
}

func (controller *ItemControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	itemCreateRequest := dtos.ItemCreateRequest{}
	helper.ReadFromRequestBody(request, &itemCreateRequest)

	userId := request.Context().Value("userId").(int)

	itemResponse := controller.ItemService.Create(request.Context(), itemCreateRequest, userId)
	dtosResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   itemResponse,
	}

	helper.WriteToResponseBody(writer, dtosResponse)
}

func (controller *ItemControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	itemUpdateRequest := dtos.ItemUpdateRequest{}
	helper.ReadFromRequestBody(request, &itemUpdateRequest)

	itemId := params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	helper.PanicIfError(err)

	itemUpdateRequest.Id = id

	userId := request.Context().Value("userId").(int)

	itemResponse := controller.ItemService.Update(request.Context(), itemUpdateRequest, userId)
	dtosResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   itemResponse,
	}

	helper.WriteToResponseBody(writer, dtosResponse)
}

func (controller *ItemControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	itemId := params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	helper.PanicIfError(err)

	controller.ItemService.Delete(request.Context(), id)
	dtosResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, dtosResponse)
}

func (controller *ItemControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	itemId := params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	helper.PanicIfError(err)

	itemResponse := controller.ItemService.FindById(request.Context(), id)
	dtosResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   itemResponse,
	}

	helper.WriteToResponseBody(writer, dtosResponse)
}

func (controller *ItemControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	itemResponses := controller.ItemService.FindAll(request.Context())
	dtosResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   itemResponses,
	}

	helper.WriteToResponseBody(writer, dtosResponse)
}

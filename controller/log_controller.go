package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kevinjuliow/dataInventarisBarang/helper"
	"github.com/kevinjuliow/dataInventarisBarang/model/dtos"
	"github.com/kevinjuliow/dataInventarisBarang/service"
)

type LogController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type LogControllerImpl struct {
	LogService service.LogService
}

func NewLogController(logService service.LogService) LogController {
	return &LogControllerImpl{
		LogService: logService,
	}
}

func (controller *LogControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logs := controller.LogService.FindAll(request.Context())

	webResponse := dtos.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   logs,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

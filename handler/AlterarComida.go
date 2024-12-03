package handler

import (
	"net/http"

	"github.com/SSPinheiro/KalCaloria/schemas"
	"github.com/gin-gonic/gin"
)

func AlterarComidaHandler(ctx *gin.Context) {
	request := AlterarComidaRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	comida := schemas.Comida{}

	if err := db.First(&comida, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "comida not found")
		return
	}

	if request.Nome != "" {
		comida.Nome = request.Nome
	}

	if request.Tipo != "" {
		comida.Tipo = request.Tipo
	}

	if request.Caloria > 0 {
		comida.Caloria = request.Caloria
	}

	if err := db.Save(&comida).Error; err != nil {
		logger.Errorf("error updating comida: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating comida")
		return
	}
	sendSucess(ctx, "update-comida", comida)
}

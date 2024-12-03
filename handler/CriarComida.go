package handler

import (
	"net/http"

	"github.com/SSPinheiro/KalCaloria/schemas"
	"github.com/gin-gonic/gin"
)

func CriarComidaHandler(ctx *gin.Context) {
	request := CriarComidaRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comida := schemas.Comida{
		Nome:    request.Nome,
		Tipo:    request.Tipo,
		Caloria: request.Caloria,
	}

	if err := db.Create(&comida).Error; err != nil {
		logger.Errorf("erro creating comida: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating comida on database")
		return
	}

	sendSucess(ctx, "create-comidaa", comida)
}

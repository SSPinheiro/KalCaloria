package handler

import (
	"net/http"

	"github.com/SSPinheiro/KalCaloria/schemas"
	"github.com/gin-gonic/gin"
)

func MostrarComidaHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	comida := schemas.Comida{}
	if err := db.First(&comida, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	sendSucess(ctx, "show-comida", comida)
}

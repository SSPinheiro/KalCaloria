package handler

import (
	"fmt"
	"net/http"

	"github.com/SSPinheiro/KalCaloria/schemas"
	"github.com/gin-gonic/gin"
)

func ExcluirComidaHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	comida := schemas.Comida{}

	if err := db.First(&comida, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	if err := db.Delete(&comida).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSucess(ctx, "delete-opening", comida)
}

package handler

import (
	"net/http"

	"github.com/SSPinheiro/KalCaloria/schemas"
	"github.com/gin-gonic/gin"
)

func ListarComidaHandler(ctx *gin.Context) {
	comidas := []schemas.Comida{}

	if err := db.Find(&comidas).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing comidas")
		sendError(ctx, http.StatusInternalServerError, "error listing comidas")
		return
	}
	sendSucess(ctx, "list-openings", comidas)
}

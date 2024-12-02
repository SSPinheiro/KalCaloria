package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarComidaHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Comida",
	})
}

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InicializarRotas(rota *gin.Engine) {
	v1 := rota.Group("/api/v1")
	{
		v1.GET("/comida", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Comida",
			})
		})
		v1.POST("/comida", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Comida",
			})
		})
		v1.DELETE("/comida", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Comida",
			})
		})
		v1.PUT("/comida", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Comida",
			})
		})
		v1.GET("/comidas", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Comidas",
			})
		})
	}
}

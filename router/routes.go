package router

import (
	"github.com/SSPinheiro/KalCaloria/handler"
	"github.com/gin-gonic/gin"
)

func InicializarRotas(rota *gin.Engine) {
	v1 := rota.Group("/api/v1")
	{
		v1.GET("/opening", handler.MostrarComidaHandler)
		v1.POST("/opening", handler.CriarComidaHandler)
		v1.DELETE("/opening", handler.ExcluirComidaHandler)
		v1.PUT("/opening", handler.AlterarComidaHandler)
		v1.GET("/openings", handler.ListarComidaHandler)
	}
}

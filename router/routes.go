package router

import (
	"github.com/SSPinheiro/KalCaloria/handler"
	"github.com/gin-gonic/gin"
)

func InicializarRotas(rota *gin.Engine) {
	handler.InitializarHandler()
	v1 := rota.Group("/api/v1")
	{
		v1.GET("/comida", handler.MostrarComidaHandler)
		v1.POST("/comida", handler.CriarComidaHandler)
		v1.DELETE("/comida", handler.ExcluirComidaHandler)
		v1.PUT("/comida", handler.AlterarComidaHandler)
		v1.GET("/comidas", handler.ListarComidaHandler)
	}
}

package router

import "github.com/gin-gonic/gin"

func Inicializar() {

	rota := gin.Default()

	InicializarRotas(rota)

	rota.Run(":8080")

}

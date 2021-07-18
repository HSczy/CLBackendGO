package main

import (
	"github.com/HSczy/CLBackendGO/utils"
)

var LogFolder = "Logs"

func main() {

	utils.Viper()
	//r := gin.Default()
	//r.Use(middlewares.Cors())
	//r.GET("/Ping", func(context *gin.Context) {
	//	context.String(http.StatusOK, "Pong")
	//})

	//r.Run(":8080")
}

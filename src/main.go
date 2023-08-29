package main

import (
	"github.com/reitmas32/receiptt-processor-technical-test/src/config"
	"github.com/reitmas32/receiptt-processor-technical-test/src/routes"

	_ "github.com/reitmas32/receiptt-processor-technical-test/src/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Receipt Processor
// @version 1.0
// @description This is a API by Technical Test by Fetch

// @contact.name Oswaldo Rafael Zamora Ramirez
// @contact.url https://github.com/reitmas32
// @contact.email rafa.zamora.rals@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:80
// @BasePath /
// @query.collection.format multi
func main() {

	config.GetRouter().GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.ReceiptsRoutes()

	config.GetRouter().Run(":" + config.PORT)
}

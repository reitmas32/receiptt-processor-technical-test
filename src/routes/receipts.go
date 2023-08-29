package routes

import (
	"github.com/reitmas32/receiptt-processor-technical-test/src/config"
	"github.com/reitmas32/receiptt-processor-technical-test/src/views"
)

func ReceiptsRoutes() {
	config.GetRouter().POST("/receipts/process", views.ReceiptsProcess)

	config.GetRouter().GET("/receipts/:id/points", views.GetPoints)
}

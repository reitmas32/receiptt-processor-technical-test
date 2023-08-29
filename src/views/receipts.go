package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/reitmas32/receiptt-processor-technical-test/src/models"
	"github.com/reitmas32/receiptt-processor-technical-test/src/services"
)

func ReceiptsProcess(c *gin.Context) {
	responseCreateReceipt := models.ResponseProcessReceipts{
		Id: "",
	}

	var receipt models.Receipt
	if err := c.ShouldBindWith(&receipt, binding.JSON); err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateReceipt)
		return
	}

	success, receipt, message := services.CreateReceipt(receipt)

	if !success {
		fmt.Println(message)
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateReceipt)
		return
	}

	responseCreateReceipt.Id = receipt.Id

	c.JSON(http.StatusOK, responseCreateReceipt)
}

func GetPoints(c *gin.Context) {
	responseGetPoints := models.ResponseGetPoints{
		Points: 0,
	}

	_, points, _ := services.GetPoints(c.Param("id"))

	responseGetPoints.Points = points

	c.JSON(http.StatusOK, responseGetPoints)
}

package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/reitmas32/receiptt-processor-technical-test/src/models"
	"github.com/reitmas32/receiptt-processor-technical-test/src/services"
)

// @Summary Create new Receipt
// @ID receipts-process
// @Tags Receipts
// @Produce json
// @Param data body models.ResponseProcessReceipts true "Schema by Create New Receipt"
// @Success 200 {object} models.ResponseProcessReceipts
// @Failure 400 {object} models.ResponseProcessReceipts
// @Router /receipts/process [post]
func ReceiptsProcess(c *gin.Context) {
	responseCreateReceipt := models.ResponseProcessReceipts{
		Id: "",
	}

	var receipt models.Receipt
	if err := c.ShouldBindWith(&receipt, binding.JSON); err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(400, responseCreateReceipt)
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

// @Summary Calculate Points
// @ID get-points
// @Tags Receipts
// @Produce json
// @Param data body models.ResponseGetPoints true "Schema by Calculate Ponits"
// @Success 200 {object} models.ResponseGetPoints
// @Failure 400 {object} models.ResponseGetPoints
// @Router /receipts/:id/points [get]
func GetPoints(c *gin.Context) {
	responseGetPoints := models.ResponseGetPoints{
		Points: 0,
	}

	success, points, _ := services.GetPoints(c.Param("id"))

	if !success {
		c.JSON(400, responseGetPoints)
	}

	responseGetPoints.Points = points

	c.JSON(http.StatusOK, responseGetPoints)
}

package views

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/reitmas32/receiptt-processor-technical-test/src/models"
	"github.com/reitmas32/receiptt-processor-technical-test/src/services"
)

// @Summary Create new Receipt
// @ID receipts-process
// @Tags Receipts
// @Produce json
// @Param data body models.Receipt true "Schema by Create New Receipt"
// @Success 200 {object} models.ResponseProcessReceipts "UUID of Receipt"
// @Failure 422 {object} map[string][]string "JSON body not is Valud missing fields"
// @Failure 400 {object} map[string]string "Error to proccess the Receipt"
// @Router /receipts/process [post]
func ReceiptsProcess(c *gin.Context) {
	responseCreateReceipt := models.ResponseProcessReceipts{
		Id: "",
	}

	var receipt models.Receipt
	if err := c.ShouldBindWith(&receipt, binding.JSON); err != nil {
		c.Header("Content-Type", "application/json")
		errorMap := map[string][]string{
			"errors": strings.Split(err.Error(), "\n"),
		}
		c.JSON(422, errorMap)
		return
	}

	success, receipt, message := services.CreateReceipt(receipt)

	if !success {
		c.Header("Content-Type", "application/json")
		errorMap := map[string]string{
			"errors": message,
		}
		c.JSON(400, errorMap)
		return
	}

	responseCreateReceipt.Id = receipt.Id

	c.JSON(http.StatusOK, responseCreateReceipt)
}

// @Summary Calculate Points
// @ID get-points
// @Tags Receipts
// @Produce json
// @Param id path string true "UUID of Receipt"
// @Success 200 {object} models.ResponseGetPoints "Total points of the Receipt"
// @Failure 404 {object} models.ResponseGetPoints "The Receipt does not exist"
// @Router /receipts/{id}/points [get]
func GetPoints(c *gin.Context) {
	responseGetPoints := models.ResponseGetPoints{
		Points: 0,
	}

	success, points, _ := services.GetPoints(c.Param("id"))

	if !success {
		c.JSON(404, responseGetPoints)
		return
	}

	responseGetPoints.Points = points

	c.JSON(http.StatusOK, responseGetPoints)
}

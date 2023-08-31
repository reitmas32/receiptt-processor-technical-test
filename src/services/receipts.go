package services

import (
	"github.com/google/uuid"
	"github.com/reitmas32/receiptt-processor-technical-test/src/config"
	"github.com/reitmas32/receiptt-processor-technical-test/src/models"
)

func CreateReceipt(receipt models.Receipt) (bool, models.Receipt, string) {

	// Note: Recomiendo hacer una verificacion de la suma de precio

	//total_expected, err := r.SumItemPrices()
	//total_real, err := strconv.ParseFloat(receipt.Total, 64)

	//if err != nil {
	//	return false, receipt, "----ERROR----- Error to convert Total or Price"
	//}

	//if total_expected != total_real {
	//
	//	return false, receipt, "----ERROR----- Sum of Items.price != Total"
	//}

	receipt.Id = uuid.New().String()

	config.DataBase = append(config.DataBase, receipt)

	return true, receipt, "Create receipt Successful"
}

func GetPoints(idReceipt string) (bool, int, string) {

	for _, r := range config.DataBase {
		if r.Id == idReceipt {
			points := r.CalculatePonts()
			return true, points, "Found successfully"
		}
	}

	return false, 0, "Not Found Receipt"
}

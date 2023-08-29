package services

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/reitmas32/receiptt-processor-technical-test/src/config"
	"github.com/reitmas32/receiptt-processor-technical-test/src/models"
)

func CreateReceipt(receipt models.Receipt) (bool, models.Receipt, string) {

	// Note: Recomiendo hacer una verificacion de la suma de precio

	//total_expected, err := sumItemPrices(receipt)
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
			points := calculatePonts(r)
			return true, points, "Found successfully"
		}
	}

	return false, 0, "Not Found Receipt"
}

func calculatePonts(receipt models.Receipt) int {
	// One point for every alphanumeric character in the retailer name.
	points := lenAlphanumerics(strings.TrimSpace(receipt.Retailer))

	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)

	// 50 points if the total is a round dollar amount with no cents.
	isInteger := (math.Mod(totalFloat, 1) == 0)

	if isInteger {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	isMultipleOfZero_25 := (math.Mod(totalFloat, 0.25) == 0)

	if isMultipleOfZero_25 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	midLen := len(receipt.Items) / 2

	fmt.Println(midLen * 5)

	points += midLen * 5

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.

	for _, item := range receipt.Items {
		isMultipleOfThree := len(strings.TrimSpace(item.ShortDescription))%3 == 0
		fmt.Println(len(item.ShortDescription))
		if isMultipleOfThree {
			priceFloat, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(priceFloat * 0.2))

		}
	}

	// 6 points if the day in the purchase date is odd.
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)

	if err != nil {
		fmt.Printf("Error to Parse purchase_date %s\n", err)
	}

	if date.Day()%2 != 0 {
		points += 6

	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)

	twoPM := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 14, 0, 0, 0, purchaseTime.Location())
	fourPM := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 16, 0, 0, 0, purchaseTime.Location())

	if err != nil {
		fmt.Printf("Error to Parse purchase_date %s\n", err)
	}

	if purchaseTime.After(twoPM) && purchaseTime.Before(fourPM) {
		points += 10

	}

	return points
}

func sumItemPrices(receipt models.Receipt) (float64, error) {
	var total float64

	for _, item := range receipt.Items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, err
		}
		total += price
	}

	return total, nil
}

func lenAlphanumerics(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

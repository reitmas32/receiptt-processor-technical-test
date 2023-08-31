package models

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/reitmas32/receiptt-processor-technical-test/src/tools"
)

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required" example:"Target"`
	PurchaseDate string `json:"purchaseDate" binding:"required" example:"2022-01-01"`
	PurchaseTime string `json:"purchaseTime" binding:"required" example:"13:01"`
	Total        string `json:"total" binding:"required" example:"1.25"`
	Items        []Item `json:"items"`
	Id           string `json:"id" example:"f985a244-11e2-4f57-a451-d7147e6a76f8"`
}

func (receipt Receipt) SumItemPrices() (float64, error) {
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

func (receipt Receipt) CalculatePonts() int {
	// One point for every alphanumeric character in the retailer name.
	points := tools.LenAlphanumerics(strings.TrimSpace(receipt.Retailer))

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

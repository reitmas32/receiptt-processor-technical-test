package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestReceipt(t *testing.T) {
	receiptJSON := `
	{
		"retailer": "Test Retailer",
		"purchaseDate": "2022-01-02",
		"purchaseTime": "08:13",
		"total": "2.65",
		"items": [
			{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
			{"shortDescription": "Dasani", "price": "1.40"}
		]
	}
	`

	var receipt Receipt
	err := json.Unmarshal([]byte(receiptJSON), &receipt)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if receipt.Retailer != "Test Retailer" {
		t.Errorf("Expected Retailer to be 'Test Retailer', but got '%s'", receipt.Retailer)
	}

	if receipt.PurchaseDate != "2022-01-02" {
		t.Errorf("Expected PurchaseDate to be '2022-01-02', but got '%s'", receipt.PurchaseDate)
	}

	if receipt.PurchaseTime != "08:13" {
		t.Errorf("Expected PurchaseTime to be '08:13', but got '%s'", receipt.PurchaseTime)
	}

	// Calculate the sum of prices from items
	var sum float64
	for _, item := range receipt.Items {
		price, _ := strconv.ParseFloat(item.Price, 64)
		sum += price
	}

	if fmt.Sprintf("%.2f", sum) != receipt.Total {
		t.Errorf("Expected Total to be '%s', but got '%s'", fmt.Sprintf("%.2f", sum), receipt.Total)
	}
}

func TestReceiptNonNegativeTotal(t *testing.T) {
	receiptJSON := `
	{
		"retailer": "Test Retailer",
		"purchaseDate": "2022-01-02",
		"purchaseTime": "08:13",
		"total": "2.65",
		"items": [
			{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
			{"shortDescription": "Dasani", "price": "1.40"}
		]
	}
	`

	var receipt Receipt
	err := json.Unmarshal([]byte(receiptJSON), &receipt)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	totalFloat, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		t.Errorf("Error to Convert Total to Float: %v", err)
	}

	if totalFloat < 0.0 {
		t.Errorf("Error The total is a negative number.: %f", totalFloat)
	}
}

func TestReceiptTotalNotFloat(t *testing.T) {
	receiptJSON := `
	{
		"retailer": "Test Retailer",
		"purchaseDate": "2022-01-02",
		"purchaseTime": "08:13",
		"total": "2.65",
		"items": [
			{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
			{"shortDescription": "Dasani", "price": "1.40"}
		]
	}
	`

	var receipt Receipt
	err := json.Unmarshal([]byte(receiptJSON), &receipt)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	_, err = strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		t.Errorf("Error to Convert Total to Float: %v", err)
	}
}

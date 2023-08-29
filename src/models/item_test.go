package models

import (
	"encoding/json"
	"testing"
)

func TestItem(t *testing.T) {
	itemJSON := `{"shortDescription": "Test Item", "price": "9.99"}`

	var item Item
	err := json.Unmarshal([]byte(itemJSON), &item)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if item.ShortDescription != "Test Item" {
		t.Errorf("Expected ShortDescription to be 'Test Item', but got '%s'", item.ShortDescription)
	}

	if item.Price != "9.99" {
		t.Errorf("Expected Price to be '9.99', but got '%s'", item.Price)
	}
}

func TestEmptyItem(t *testing.T) {
	itemJSON := `{}`

	var item Item
	err := json.Unmarshal([]byte(itemJSON), &item)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if item.ShortDescription != "" {
		t.Errorf("Expected ShortDescription to be empty, but got '%s'", item.ShortDescription)
	}

	if item.Price != "" {
		t.Errorf("Expected Price to be empty, but got '%s'", item.Price)
	}
}

func TestInvalidJSON(t *testing.T) {
	itemJSON := `invalid JSON`

	var item Item
	err := json.Unmarshal([]byte(itemJSON), &item)
	if err == nil {
		t.Error("Expected error while unmarshaling invalid JSON, but got none")
	}
}

func TestInvalidType(t *testing.T) {
	itemJSON := `{"shortDescription": 123, "price": "9.99"}`

	var item Item
	err := json.Unmarshal([]byte(itemJSON), &item)
	if err == nil {
		t.Error("Expected error due to invalid type, but got none")
	}
}

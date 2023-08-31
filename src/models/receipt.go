package models

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required" example:"Target"`
	PurchaseDate string `json:"purchaseDate" binding:"required" example:"2022-01-01"`
	PurchaseTime string `json:"purchaseTime" binding:"required" example:"13:01"`
	Total        string `json:"total" binding:"required" example:"1.25"`
	Items        []Item `json:"items"`
	Id           string `json:"id" example:"f985a244-11e2-4f57-a451-d7147e6a76f8"`
}

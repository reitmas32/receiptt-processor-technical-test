package models

type Item struct {
	ShortDescription string `json:"shortDescription" example:"Pepsi - 12-oz"`
	Price            string `json:"price" example:"1.25"`
}

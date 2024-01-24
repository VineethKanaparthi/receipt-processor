package model

import (
	"errors"
	"time"
)

// Receipt represents the structure of a receipt.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items" binding:"dive"`
	Total        string `json:"total" binding:"numeric"`
}

// Validate validates that the receipt variables are in the correct format
func (receipt *Receipt) Validate() error {
	if receipt.PurchaseDate != "" {
		_, err := time.Parse("2006-01-02", receipt.PurchaseDate)
		if err != nil {
			return errors.New("field `purchaseDate` is not in the correct format")
		}
	}

	if receipt.PurchaseTime != "" {
		_, err := time.Parse("15:04", receipt.PurchaseTime)
		if err != nil {
			return errors.New("field `purchaseTime` is not in the correct format")
		}
	}
	return nil
}

// Item represents the structure of an item in a receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price" binding:"numeric"`
}

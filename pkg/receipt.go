package model

// Receipt represents the structure of a receipt.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents the structure of an item in a receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

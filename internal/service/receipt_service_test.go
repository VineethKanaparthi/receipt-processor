package service

import (
	"testing"

	model "github.com/VineethKanaparthi/receipt-processor/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	// Test calculatePoints function

	tests := []struct {
		receipt model.Receipt
		points  int
	}{
		{
			model.Receipt{
				Retailer:     "",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			22,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			28,
		},
		{
			model.Receipt{
				Retailer:     "Target 1",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			29,
		},
		{
			model.Receipt{
				Retailer:     "Target$",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			28,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.25",
			},
			53,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.00",
			},
			103,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
					{
						ShortDescription: "Doritos Nacho Cheese 18 FL OZ",
						Price:            "25.00",
					},
				},
				Total: "35.35",
			},
			33,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			22,
		},

		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:59",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			28,
		},

		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "14:00",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			28,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "14:01",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			38,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "15:59",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			38,
		},
		{
			model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "16:00",
				Items: []model.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			28,
		},
	}
	for _, test := range tests {
		t.Run("CalculatePoints", func(t *testing.T) {
			points := CalculatePoints(&test.receipt)

			// Add assertions based on the expected points for this receipt
			assert.Equal(t, test.points, points)
		})
	}

}

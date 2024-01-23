package service

import (
	model "github.com/VineethKanaparthi/receipt-processor/pkg"
	"github.com/google/uuid"
)

func ProcessReceipt(receipt *model.Receipt) (string, error) {
	return uuid.New().String(), nil
}

func GetPoints(id string) (int64, error) {
	return 32, nil
}

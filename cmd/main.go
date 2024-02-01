package main

import (
	"github.com/VineethKanaparthi/receipt-processor/internal/database"
	"github.com/VineethKanaparthi/receipt-processor/internal/server"
)

// main function initializes and runs the server
func main() {
	server := server.NewReceiptServer()
	db := database.NewBoltDatabase("receipts.db")
	server.DB = db
	server.Run(":8080")
	defer db.Close()
}

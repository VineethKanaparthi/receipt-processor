package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/VineethKanaparthi/receipt-processor/internal/service"
	model "github.com/VineethKanaparthi/receipt-processor/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

func ProcessReceipt(c *gin.Context, db *bolt.DB) {
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if err := receipt.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id, err := service.ProcessReceipt(&receipt, db)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process the receipt, please try again"})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": id})
		}
	}
}

func GetPoints(c *gin.Context, db *bolt.DB) {
	id := c.Params.ByName("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a uuid"})
	} else {
		points, err := service.GetPoints(id, db)
		if err != nil {
			if errors.Is(err, service.ErrIdNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get points for the id"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"points": points})
		}
	}
}

package router

import (
	"net/http"

	"github.com/VineethKanaparthi/receipt-processor/internal/service"
	model "github.com/VineethKanaparthi/receipt-processor/pkg"
	"github.com/gin-gonic/gin"
)

// TODO:
// documentation
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// TODO:
	// 1) documentation
	// 2) logging?
	r.POST("/receipts/process", func(c *gin.Context) {
		var receipt model.Receipt
		if err := c.BindJSON(&receipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format"})
		} else {
			id, err := service.ProcessReceipt(&receipt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process the receipt, please try again"})
			} else {
				c.JSON(http.StatusOK, gin.H{"id": id})
			}
		}

	})

	// TODO:
	// 1) documentation
	// 2) logging?
	r.GET("receipts/:id/points", func(c *gin.Context) {
		id := c.Params.ByName("id")
		points, err := service.GetPoints(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process the receipt, please try again"})
		} else {
			c.JSON(http.StatusOK, gin.H{"points": points})
		}

	})
	return r

}

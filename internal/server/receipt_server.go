package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/VineethKanaparthi/receipt-processor/internal/service"
	model "github.com/VineethKanaparthi/receipt-processor/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type ReceiptServer struct {
	DB *bolt.DB
	*gin.Engine
}

type ReceiptResponse struct {
	ID string `json:"id"`
}

// NewBoltDatabase initializes the database
func NewBoltDatabase(dbname string) *bolt.DB {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the points bucket in the database
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("points"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// CloseDB closes the database.
func (rs *ReceiptServer) CloseDB() {
	if rs.DB != nil {
		if err := rs.DB.Close(); err != nil {
			log.Println("Error closing database:", err)
		}
	}
}

// NewReceiptServer initializes the server, creates a database with dbname and sets up the router
func NewReceiptServer(dbname string) *ReceiptServer {
	rs := &ReceiptServer{}

	router := gin.Default()
	// POST /receipts/process endpoint
	router.POST("/receipts/process", rs.processReceipt)
	// GET /receipts/:id/points endpoint
	router.GET("receipts/:id/points", rs.getPoints)

	rs.Engine = router
	rs.DB = NewBoltDatabase(dbname)
	return rs
}

func (rs *ReceiptServer) processReceipt(c *gin.Context) {
	var receipt model.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := receipt.Validate(); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := service.ProcessReceipt(&receipt, rs.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process the receipt, please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (rs *ReceiptServer) getPoints(c *gin.Context) {
	id := c.Params.ByName("id")
	if _, err := uuid.Parse(id); err != nil {
		handleError(c, http.StatusBadRequest, "id is not a uuid")
		return
	}

	points, err := service.GetPoints(id, rs.DB)
	if err != nil {
		handleGetPointsError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

func handleGetPointsError(err error, c *gin.Context) {
	if errors.Is(err, service.ErrIdNotFound) {
		handleError(c, http.StatusNotFound, err.Error())
	} else {
		handleError(c, http.StatusInternalServerError, "failed to get points for the id")
	}
}

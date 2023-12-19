package main

import (
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

type Customer struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db.AutoMigrate(&Customer{})
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, ITMX!"})
	})

	r.POST("/customers", createCustomer)
	r.PUT("/customers/:id", updateCustomer)
	r.DELETE("/customers/:id", deleteCustomer)
	r.GET("/customers/:id", getCustomer)
	r.POST("/initData", initializeData)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&customer)
	c.JSON(http.StatusCreated, customer)
}

func updateCustomer(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var updatedCustomer Customer

	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Model(&Customer{}).Where("id = ?", id).Updates(updatedCustomer)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, updatedCustomer)
}

func deleteCustomer(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	db.Delete(&Customer{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

func getCustomer(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var customer Customer
	result := db.First(&customer, id)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func initializeData(c *gin.Context) {
	db.Create(&Customer{Name: "Kitsada ", Age: 30})
	db.Create(&Customer{Name: "suparut", Age: 25})
	c.JSON(http.StatusOK, gin.H{"message": "Data initialized"})
}

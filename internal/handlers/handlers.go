package handlers

import (
	"net/http"
	"strconv"

	"book.com/internal/db"
	"book.com/internal/models"
	"book.com/internal/storage"

	"github.com/gin-gonic/gin"
)

func UploadBook(c *gin.Context) {
	name := c.PostForm("name")
	author := c.PostForm("author")
	priceStr := c.PostForm("price")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	blobName, err := storage.UploadFile(file, header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Azure upload failed"})
		return
	}

	book := models.Book{
		Name:    name,
		Author:  author,
		Price:   price,
		FileURL: blobName,
	}

	if err := db.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database save failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book uploaded", "book": book})
}

func SearchBook(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query required"})
		return
	}

	var books []models.Book
	if err := db.DB.Where("name ILIKE ?", "%"+q+"%").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	var results []gin.H
	for _, b := range books {
		url := storage.GenerateBlobSAS(b.FileURL)

		results = append(results, gin.H{
			"name":   b.Name,
			"author": b.Author,
			"price":  b.Price,
			"url":    url,
		})
	}

	c.JSON(http.StatusOK, results)
}

package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-crud/models"
	"net/http"
)

var books = []models.Book{
	{ID: "1", Title: "Go Programming", Author: "John"},
	{ID: "2", Title: "REST API", Author: "Alice"},
}

// Получить все книги
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Получить одну книгу
func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Добавить новую книгу
func createBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Обновить книгу
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Удалить книгу
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run("localhost:8082")
}

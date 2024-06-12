package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	{ID: "4", Title: "Moby Dick", Author: "Herman Melville", Quantity: 3},
	{ID: "5", Title: "The Odyssey", Author: "Homer", Quantity: 4},
	{ID: "6", Title: "Ulysses", Author: "James Joyce", Quantity: 2},
	{ID: "7", Title: "Pride and Prejudice", Author: "Jane Austen", Quantity: 7},
	{ID: "8", Title: "The Divine Comedy", Author: "Dante Alighieri", Quantity: 1},
	{ID: "9", Title: "The Brothers Karamazov", Author: "Fyodor Dostoevsky", Quantity: 3},
	{ID: "10", Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Quantity: 5},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, books)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			if book.Quantity > 0 {
				books[i].Quantity--
				c.IndentedJSON(http.StatusOK, books)
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "book out of stock"})
				return
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func returnBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books[i].Quantity++
			c.IndentedJSON(http.StatusOK, books)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.DELETE("/books/:id", deleteBook)
	router.POST("/books", createBook)
	router.PUT("/books/:id/checkout", checkoutBook)
	router.PUT("/books/:id/return", returnBook)
	router.Run(":8080")
}

package main

import (
	"errors"
	"net/http"
	"strconv"

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
}

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func findByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := findByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func buyBook(c *gin.Context) {
	id, okID := c.GetQuery("id")
	buyQuantity, okQuantity := c.GetQuery("quantity")
	buyQuantityInt, _ := strconv.Atoi(buyQuantity)

	if !okID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id parameter."})
		return
	}

	if !okQuantity {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing quantity parameter."})
		return
	}

	book, err := findByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity < buyQuantityInt {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not enough books in store"})
		return
	}

	book.Quantity -= buyQuantityInt
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookByID)
	router.PATCH("/books/buy", buyBook)
	router.Run("localhost:8080")
}

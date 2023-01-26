package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//1. go run main.go
//2. in separate terminal send curl request from comments in main()

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
	//To see all books
	//curl 'localhost:8080/books'
	router.GET("/books", getAllBooks)
	//Add book in library
	//curl 'localhost:8080/books' -i -H "Content-Type: application/json" -d "@body.json" --request "POST"
	router.POST("/books", createBook)
	//To see Book by id
	//curl 'localhost:8080/books/3'
	router.GET("/books/:id", getBookByID)
	//To buy 3 Books by id
	//curl 'localhost:8080/books/buy?id=3&quantity=3' --request "PATCH"
	router.PATCH("/books/buy", buyBook)
	router.Run("localhost:8080")
}

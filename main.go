package main

import (
	// "errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct{
  ID string `json:"id"`
  Title string  `json:"title"`
  Author string `json:"author"`
}

var books =[]book{
  {ID : "1", Title: "Chainsaw Man", Author: "Tatsuki Fujimoto"},
  {ID : "2", Title: "One Punch Man", Author: "One"},
  {ID : "3", Title: "One Piece", Author: "Eiichiro Oda"},
}

func getBooks(c *gin.Context){
  c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusFound, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Book not found"})
}

func main(){
  router := gin.Default()
  router.GET("/books", getBooks)
  router.GET("/books", getBookByID)
  router.POST("/books", addBook)
  
  
  router.Run("localhost:1111")
}

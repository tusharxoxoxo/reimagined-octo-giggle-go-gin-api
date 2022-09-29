package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct{
  ID string `json:"id"`
  Title string  `json:"title"`
  Author string `json:"author"`
  Quantity int `json:"quantity"`
}

var books =[]book{
  {ID : "1", Title: "Chainsaw Man", Author: "Tatsuki Fujimoto", Quantity: 5},
  {ID : "2", Title: "One Punch Man", Author: "One", Quantity: 1},
  {ID : "3", Title: "One Piece", Author: "Eiichiro Oda", Quantity: 1055},
}

func getBooks(c *gin.Context){
  c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	
}

func main(){
  router := gin.Default()
  router.GET("/books", getBooks)
  router.Run("localhost:8080")
}

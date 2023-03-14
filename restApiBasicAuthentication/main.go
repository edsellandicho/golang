package main

//net/http is a GO package that provides HTTP client and server implementation.
//   Get, Head, Post, and PostForm make HTTP (or HTTPS) requests.
//github.com/gin-gonic/gin is a GO Module package for using Gin. For your reference, Gin is a web framework written in Go.
//   It features a martini-like API with performance that is up to 40 times faster thanks to httprouter.
//   If you need performance and good productivity, you will love Gin.
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**********************************************************
* TEMPORARY DATA STORED IN MEMORY FOR TESTING PURPOSES
* In real world, this is usually a database.
***********************************************************/
// item represents the data about a grocery item.
type item struct {
	ID       string  `json:"id"`
	ItemName string  `json:"itemname"`
	Aisle    uint    `json:"aisle"`
	InStock  bool    `json:"instock"`
	Price    float64 `json:"price"`
}

// items is a collection of grocery item, it is set as a slice data type (abstraction of an array).
var items = []item{
	{ID: "1", ItemName: "Fruit Snack", Aisle: 5, InStock: true, Price: 3.99},
	{ID: "2", ItemName: "Z Bar", Aisle: 3, InStock: false, Price: 4.25},
	{ID: "3", ItemName: "Dinner Roll", Aisle: 7, InStock: true, Price: 5.50},
}

func main() {
	router := gin.Default()
	//add basic authentication security in gin framework - for testing purposes
	accounts := gin.Accounts{"edsel": "testpassword"}
	router.Use(gin.BasicAuth(accounts))
	//end: adding basic authentication
	router.GET("/items", getItems) //Test login service
	router.Run("localhost:8080")
}

// GET ITEMS
// To test: curl -u edsel:testpassword http://localhost:8080/items
// TO test: curl -H 'Authorization:Basic ZWRzZWw6dGVzdHBhc3N3b3Jk' http://localhost:8080/items
func getItems(ginContext *gin.Context) {
	fmt.Println("Test Login requested")
	ginContext.IndentedJSON(http.StatusOK, items)
}

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
	"golang.org/x/exp/slices"
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

	router.POST("/items", postItem)             //Create or Add a new item
	router.GET("/items", getItems)              //Search or GET multiple items in return
	router.GET("/items/:id", getItemByID)       //Retrieve or Get an item
	router.PATCH("/items/:id", updateItemByID)  //Update or Edit an item..we use PATCH instead of PUT.
	router.DELETE("/items/:id", deleteItemByID) //Delete or Remove an item

	router.Run("localhost:8080")
}

// RETRIEVE or GET multiple items.
// getItems return with the list of all items in JSON format.
// To test: curl http://localhost:8080/items
func getItems(ginContext *gin.Context) {
	fmt.Println("RETRIEVE all items requested...")
	ginContext.IndentedJSON(http.StatusOK, items)
}

// RETRIEVE or GET an item.
// getItemByID locates the item whose ID value matches the id
// parameter sent by the client, then returns that item as a response.
// To test: curl http://localhost:8080/items/2
func getItemByID(ginContext *gin.Context) {
	fmt.Println("RETRIEVE an item requested...")
	id := ginContext.Param("id")

	// Loop through the list of items, looking for
	// an item whose ID value matches the parameter.
	for _, eachItem := range items {
		if eachItem.ID == id {
			ginContext.IndentedJSON(http.StatusOK, eachItem)
			return
		}
	}
	ginContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// CREATE or ADD a new item.
// postItem adds a new item from the request body, in JSON format.
// To test: curl -X POST -H "Content-Type: application/json" -d '{"id": "4","itemname": "Milk","aisle": 8,"instock": true,"price": 7.05}' http://localhost:8080/items
func postItem(ginContext *gin.Context) {
	fmt.Println("CREATE item requested...")
	var newItem item

	// Call BindJSON to bind the received JSON to
	// newItem.
	if err := ginContext.BindJSON(&newItem); err != nil {
		return
	}

	// Add the new item to the slice data type (abstraction of an array).
	items = append(items, newItem)
	ginContext.IndentedJSON(http.StatusCreated, newItem)
}

// UPDATE an item.
// updateItemByID locates the item whose ID value matches the id
// parameter sent by the client, then update the item from the items slice daya type.
// Note that we are using PATCH as http method instead of PUT...its a GO thing for UPDATE :-)
// To test: curl -X PATCH -H "Content-Type: application/json" -d '{"id": "2","itemname": "brioche bread","aisle": 9, "price": 4.35}' http://localhost:8080/items/2
func updateItemByID(ginContext *gin.Context) {
	fmt.Println("UPDATE item requested...")
	id := ginContext.Param("id")
	var newItem item

	// Call BindJSON to bind the received JSON to
	// newItem.
	if err := ginContext.BindJSON(&newItem); err != nil {
		return
	}

	// Loop through the list of items, looking for
	// an item whose ID value matches the parameter.
	for counter, eachItem := range items {
		if eachItem.ID == id {
			fmt.Println("UPDATE this item: ", eachItem)
			fmt.Println("with this new item: ", newItem)
			items[counter] = newItem
			ginContext.IndentedJSON(http.StatusOK, eachItem)
			return
		}
	}
	ginContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// DELETE an item.
// deleteItemByID locates the item whose ID value matches the id
// parameter sent by the client, then delete or remove it from the items slice daya type.
// To test: curl -X DELETE http://localhost:8080/items/2
func deleteItemByID(ginContext *gin.Context) {
	fmt.Println("DELETE item requested...")
	id := ginContext.Param("id")
	// Loop through the list of items, looking for
	// an item whose ID value matches the parameter.
	for counter, eachItem := range items {
		if eachItem.ID == id {
			fmt.Printf("Found the item with ID %v for deletion...\n", id)
			fmt.Println(eachItem)
			fmt.Printf("Removing index number %d \n", counter)
			items = slices.Delete(items, counter, counter+1)
			ginContext.IndentedJSON(http.StatusOK, eachItem)
			fmt.Println("Remaining items:")
			fmt.Println(items)
			return
		}
	}
	ginContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/base64"
)

type Item struct {
	ID       string  `json:"id"`
	ItemName string  `json:"itemname"`
	Aisle    uint    `json:"aisle"`
	InStock  bool    `json:"instock"`
	Price    float64 `json:"price"`
}

func main() {
	fmt.Println("CONSUME A REST API or CLIENT REST API")
	//getIemByID()
	getItems()
	/*
		deleteItem()
		getItems()
		createNewItem()
		getItems()
		updateExistingItem()
		getItems()
	*/
	fmt.Println("*** END ***")

}

func getItems() {
	fmt.Println("********************************")
	fmt.Println("RETRIEVE ALL ITEMS (HTTP GET)")
	fmt.Println("********************************")
	//Create an HTTP Client (so you can control over the http headers and other settings)
	client := &http.Client{}
	//create a new http request
	request, err := http.NewRequest("GET", "http://localhost:8080/items", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	//set the http headers
	fmt.Println("Set http header")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	//********************************
	//How to add basic authentication
	//********************************
	fmt.Println("Set basic authenticaion...")
	user := "edsel"
	password := "testpassword"
	basicAuthenticationHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	fmt.Println("!" + basicAuthenticationHeaderValue + "!")
	request.Header.Set("Authorization", basicAuthenticationHeaderValue)
	//****END: basic aunthentication *************

	fmt.Println("Sending request to REST API...")
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("REST Response: ", response)

	fmt.Println("Closing the connection")
	defer response.Body.Close() //close the response body when finished.

	fmt.Println("Converting the response body into json format")
	bodyBytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println("bodyBytes value:")
	//fmt.Println(bodyBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Converting the json format into go structure data type using unmarshal")
	var responseObject []Item
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println("*******************")
	fmt.Println("REST API Response:")
	fmt.Println("*******************")
	fmt.Println(responseObject)
}

func getIemByID() {
	fmt.Println("*****************************")
	fmt.Println("RETRIEVE AN ITEM (HTTP GET)")
	fmt.Println("*****************************")
	//Create an HTTP Client (so you can control over the http headers and other settings)
	client := &http.Client{}
	//create a new http request
	request, err := http.NewRequest("GET", "http://localhost:8080/items/1", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	//set the http headers
	fmt.Println("Set http header")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	fmt.Println("Sending request to REST API...")
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("REST Response: ", response)

	fmt.Println("Closing the connection")
	defer response.Body.Close() //close the response body when finished.

	fmt.Println("Converting the response body into json format")
	bodyBytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println("bodyBytes value:")
	//fmt.Println(bodyBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Converting the json format into go structure data type using unmarshal")
	var responseObject Item
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println("*******************")
	fmt.Println("REST API Response:")
	fmt.Println("*******************")
	fmt.Println(responseObject)

}

func createNewItem() {
	fmt.Println("*****************************")
	fmt.Println("CREATE NEW ITEM (HTTP POST)")
	fmt.Println("*****************************")
	item := Item{"5", "Goldfish crackers", 12, false, 1.52}
	fmt.Println("json marshall (Convert json format into byte array format before sending to HTTP request)")
	jsonRequestData, err := json.Marshal(item)
	if err != nil {
		fmt.Print(err.Error())
	}

	//Create an HTTP Client (so you can control over the http headers and other settings)
	client := &http.Client{}
	//create a new http request
	request, err := http.NewRequest("POST", "http://localhost:8080/items", bytes.NewBuffer(jsonRequestData))
	if err != nil {
		fmt.Print(err.Error())
	}
	//set the http headers
	fmt.Println("Set http header")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	fmt.Println("Sending request to REST API...")
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("REST Response: ", response)

	fmt.Println("Closing the connection")
	defer response.Body.Close() //close the response body when finished.

	fmt.Println("Converting the response body into json format")
	bodyBytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println("bodyBytes value:")
	//fmt.Println(bodyBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Converting the json format into go structure data type using unmarshal")
	var responseObject Item
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println("*******************")
	fmt.Println("REST API Response:")
	fmt.Println("*******************")
	fmt.Println(responseObject)
}

func updateExistingItem() {
	fmt.Println("**************************************************")
	fmt.Println("UPDATE EXISTING ITEM (HTTP PATCH instead of PUT)")
	fmt.Println("**************************************************")
	item := Item{"2", "brioche bread", 15, true, 7.35}
	fmt.Println("json marshall (Convert json format into byte array format before sending to HTTP request)")
	jsonRequestData, err := json.Marshal(item)
	if err != nil {
		fmt.Print(err.Error())
	}

	//Create an HTTP Client (so you can control over the http headers and other settings)
	client := &http.Client{}
	//create a new http request
	request, err := http.NewRequest("PATCH", "http://localhost:8080/items/2", bytes.NewBuffer(jsonRequestData))
	if err != nil {
		fmt.Print(err.Error())
	}
	//set the http headers
	fmt.Println("Set http header")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	fmt.Println("Sending request to REST API...")
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("REST Response: ", response)

	fmt.Println("Closing the connection")
	defer response.Body.Close() //close the response body when finished.

	fmt.Println("Converting the response body into json format")
	bodyBytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println("bodyBytes value:")
	//fmt.Println(bodyBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Converting the json format into go structure data type using unmarshal")
	var responseObject Item
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println("*******************")
	fmt.Println("REST API Response:")
	fmt.Println("*******************")
	fmt.Println(responseObject)
}

func deleteItem() {
	fmt.Println("*****************************")
	fmt.Println("DELETE AN ITEM (HTTP DELETE)")
	fmt.Println("*****************************")
	//Create an HTTP Client (so you can control over the http headers and other settings)
	client := &http.Client{}
	//create a new http request
	request, err := http.NewRequest("DELETE", "http://localhost:8080/items/1", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	//set the http headers
	fmt.Println("Set http header")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	fmt.Println("Sending request to REST API...")
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("REST Response: ", response)

	fmt.Println("Closing the connection")
	defer response.Body.Close() //close the response body when finished.

	fmt.Println("Converting the response body into json format")
	bodyBytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println("bodyBytes value:")
	//fmt.Println(bodyBytes)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Converting the json format into go structure data type using unmarshal")
	var responseObject Item
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println("*******************")
	fmt.Println("REST API Response:")
	fmt.Println("*******************")
	fmt.Println(responseObject)
}

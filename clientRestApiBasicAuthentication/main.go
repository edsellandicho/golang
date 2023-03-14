package main

import (
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
	fmt.Println("CONSUME A REST API USING BASIC AUTHENTICATION")
	getItems()
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

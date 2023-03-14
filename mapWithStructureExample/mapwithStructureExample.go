package main

import "fmt"

func main() {
	fmt.Println("******************************")
	fmt.Println("* MAP with Structure Example *")
	fmt.Println("******************************")
	//initialize using make
	var userActives = make(map[string]bool)
	userActives["eds"] = true
	userActives["ja"] = false

	//add
	userActives["jaz"] = true

	//read all
	fmt.Println(userActives)
	//read value per key
	fmt.Println(userActives["ja"])

	//loop all map key value pairs
	for key, value := range userActives {
		fmt.Printf("Key %v | Value: %v \n", key, value)
	}

	//Let's use Structure
	fmt.Println("****************************************************")
	fmt.Println("MAP with Structure Example")
	fmt.Println("to support value with different or mixed data type")
	fmt.Println("****************************************************")

	//first character must be capital
	type UserData struct {
		firstName          string
		lastName           string
		email              string
		totalNumberOfSeats uint
		isValidUser        bool
	}

	//how to use structure
	var userData = UserData{firstName: "Edsel",
		lastName:           "Landicho",
		email:              "test@mail.com",
		totalNumberOfSeats: 5,
		isValidUser:        true,
	}

	//let's incorporate in your map key value pairs
	var userInformations = make(map[string]UserData)
	userInformations["edsel"] = userData
	userInformations["ja"] = UserData{firstName: "Ja",
		lastName:           "Landicho",
		email:              "ja@mail.com",
		totalNumberOfSeats: 8,
		isValidUser:        false,
	}
	fmt.Println(userInformations)

}

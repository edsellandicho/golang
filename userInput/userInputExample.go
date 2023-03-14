package main

import "fmt"

func main() {
	fmt.Println("**********************")
	fmt.Println("* User Input Example *")
	fmt.Println("***********************")
	var firstName = "Edsel"
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Your first name is ", firstName)

}

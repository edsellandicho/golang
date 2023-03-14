package main

import "fmt"

func main() {
	fmt.Println("*****************************")
	fmt.Println("* IF ELSE STATEMENT Example *")
	fmt.Println("*****************************")
	fmt.Println("variable: age is 45")
	fmt.Println("if age is greater than 65 then run the if block statement.")
	fmt.Println("otherwise, print the else or otherwise block statement.")
	var age int = 45
	if age > 65 {
		//statement to execute if condition is true
		fmt.Println("Result: if block statement")
	} else {
		fmt.Println("Result: else or otherwise block statement")
	}

}

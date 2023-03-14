package main

import "fmt"

func main() {
	fmt.Println("***************************************")
	fmt.Println("* IF ELSE IF LADDER STATEMENT Example *")
	fmt.Println("***************************************")
	fmt.Println("variable: ageOfMommy is 42")
	var ageOfMommy int = 42
	if ageOfMommy == 38 {
		//statement to execute if condition is true
		fmt.Println("Result: if block statement - ageOfMommy is 38")
	} else if ageOfMommy == 42 {
		//statement to execute if condition is true
		fmt.Println("Result: else if block statement - ageOfMommy is 42")
	} else {
		fmt.Println("Result: otherwise block statement. ageOfMommy is not 38 and not 42")
	}
}

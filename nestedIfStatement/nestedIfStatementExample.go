package main

import "fmt"

func main() {
	fmt.Println("*******************************")
	fmt.Println("* NESTED IF STATEMENT Example *")
	fmt.Println("*******************************")
	fmt.Println("variable: ageOfMommy is 42")
	fmt.Println("variable: ageOfDaddy is 44")
	var ageOfMommy int = 42
	var ageOfDaddy int = 44
	if ageOfMommy == 42 {
		//statement to execute if condition is true
		fmt.Println("Result: if block statement - ageOfMommy is 42")
		if ageOfDaddy == 44 {
			//statement to execute if condition is true
			fmt.Println("Result: nested if block statement - ageOfDaddy is 44")
		}
	}
}

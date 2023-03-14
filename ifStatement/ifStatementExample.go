package main

import "fmt"

func main() {
	fmt.Println("************************")
	fmt.Println("* IF STATEMENT Example *")
	fmt.Println("************************")
	fmt.Println("variable: allow is true")
	var allow bool = true
	if allow {
		//statement to execute if condition is true
		fmt.Println("Result: If statement block statement")
	}
}

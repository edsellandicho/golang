package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

var ApplicationCertificationNumber = "ABCD1234"

func Cook(dish string) {
	fmt.Printf("Cooking %v in progress\n", dish)
	fmt.Printf("Cooking %v is done.\n", dish)
}

func drink(beverage string) string {
	fmt.Printf("Drinking %v in progress\n", beverage)
	fmt.Printf("Drinking %v is done.\n", beverage)
	return "success"
}

func washDishes(washableItems string) {
	fmt.Printf("Washing %v in progress\n", washableItems)
	fmt.Printf("Washing %v is done.\n", washableItems)
}

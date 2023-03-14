package internal

import "fmt"

var ApplicationCertificationNumber = "ABCD1234"

func Cook(dish string) {
	fmt.Printf("Cooking %v in progress\n", dish)
	fmt.Printf("Cooking %v is done.\n", dish)
}

func Drink(beverage string) string {
	fmt.Printf("Drinking %v in progress\n", beverage)
	fmt.Printf("Drinking %v is done.\n", beverage)
	return "success"
}

func washDishes(washableItems string) {
	fmt.Printf("Washing %v in progress\n", washableItems)
	fmt.Printf("Washing %v is done.\n", washableItems)
}

/* func main() {
	cook("bacon")
	drink("juice")
	washDishes("glasses, plates and silverwares")
} */

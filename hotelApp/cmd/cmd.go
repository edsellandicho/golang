package main

import (
	"fmt"

	"edsellandicho.com/internal"
)

func main() {
	fmt.Println("Test cmd")
	internal.Cook("egg")
	a := internal.Drink("juice")
	fmt.Println(a)

}

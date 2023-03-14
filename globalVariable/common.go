package main

import "fmt"

var ApplicationName = "Use Global Variable value - successful!"

func washing(washableItems string) {
	fmt.Printf("Washing %v in progress\n", washableItems)
	fmt.Printf("Washing %v is done.\n", washableItems)
}

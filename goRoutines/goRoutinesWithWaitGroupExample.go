package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{} //goroutine step 1: Need to declare as package level (global) variable so its visible to all functions

func main() {
	fmt.Println("*************************************")
	fmt.Println("* GOROUTINES WITH WAITGROUP EXAMPLE *")
	fmt.Println("*************************************")

	fmt.Println("WaitGroup variable: set in package level")
	fmt.Println("WaitGroup.add: Setting to 1 goroutine/thread...")
	waitGroup.Add(2) //goroutine step 2
	fmt.Println("Start processing cooking function as new goroutine/thread...")
	go cooking("egg")
	fmt.Println("Start processing drinking function...")
	go drinking("coffee")

	fmt.Println("WaitGroup.Wait status: In progress...until all goroutines/threads are completed or released.")
	waitGroup.Wait() //wait until all goroutines/threads are released
	fmt.Println("WaitGroup.Wait status: DONE")

	fmt.Println("*** PROGRAM WITH GOROUTINE COMPLETED ***")

}

func cooking(dish string) {
	fmt.Printf("Cooking %v in progress...\n", dish)
	fmt.Printf("Cooking %v is done.\n", dish)
	fmt.Println("Cooking WaitGroup.done status: in progress...")
	waitGroup.Done()
	fmt.Println("Cooking WaitGroup done status: completed")
}

func drinking(beverage string) {
	fmt.Printf("Drinking %v in progress\n", beverage)
	fmt.Printf("Drinking %v is done.\n", beverage)
	fmt.Println("Drinking WaitGroup.done status: in progress...")
	waitGroup.Done()
	fmt.Println("Drinking WaitGroup done status: completed")
}

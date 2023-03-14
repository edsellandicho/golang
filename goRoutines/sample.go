/* package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{} //goroutine step 1

func main() {
	fmt.Println("Application start")
	waitGroup.Add(1) //goroutine step 2
	go functionToProcess()
	waitGroup.Wait() //groutine step 4
	fmt.Println("Application completed")
}

func functionToProcess() {
	fmt.Println("functionToProcess block statement")
	waitGroup.Done() //groutine step 3
}
*/
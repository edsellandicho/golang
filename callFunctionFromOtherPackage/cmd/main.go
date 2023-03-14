package main

import (
   "fmt"
   "edsellandicho.com/sampleapp/internal"
)

func main() {
   fmt.Println("Call function from the other package sample project")
   internal.TurnLeft()
   internal.TurnRight()
}

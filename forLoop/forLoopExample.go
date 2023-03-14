package main

import "fmt"

func main() {
	fmt.Println("********************")
	fmt.Println("* GOLANG *")
	fmt.Println("* FOR LOOP EXAMPLE *")
	fmt.Println("********************")
	fmt.Println("")

	fmt.Println("Use Case #1: Standard way")
	for i := 1; i <= 3; i++ {
		fmt.Println("for loop block statement run...")
	}

	fmt.Println("")
	fmt.Print("Use Case #2: Initialization is outside for loop statement")
	fmt.Println(" and postcondtion is inside the for loop statement. ")
	fmt.Println(" For loop using int data type.")

	var i int = 1
	for i <= 3 {
		fmt.Println("for loop block statement run...")
		i++
	}

	fmt.Println("")
	fmt.Println("Use Case #3: loop using boolean as condition")
	i = 1
	var isContinue bool = true
	for isContinue {
		fmt.Println("for loop block statement run...")
		if i == 3 {
			isContinue = false
		}
		i++
	}

	fmt.Println("")
	fmt.Println("Use Case #4: Infinite Loop: The only way to stop is by using the break command within the loop block statement")
	i = 1
	for {
		fmt.Println("for loop run...")
		if i == 3 {
			fmt.Println("running break command...to stop")
			break
		}
		i++
	}

	fmt.Println("*********************************")
	fmt.Println("USING RANGE KEYWORD")
	fmt.Println("TO LOOP A MAP, ARRAY, STRING")
	fmt.Println("*********************************")
	fmt.Println("")

	fmt.Println("Use Case #5: loop a string using range keyword. Or iterate over each character in a string.")
	for i, j := range "Hello Edsel" {
		fmt.Printf("index %v | character: %s \n", i, string(j))
		//fmt.Print("index number: ", i)
		//fmt.Println("; character: ", string(j))
	}

	fmt.Println("")
	fmt.Println("Use Case #6: loop a map using range keyword. Where map is a key and value pairs.")
	mapSample := map[int]string{
		44: "Edsel",
		42: "Jan",
		7:  "Jaz",
		4:  "Jil",
	}

	fmt.Println("Printing the map using keys only")
	fmt.Println("Syntax: for key := range map {")
	for key := range mapSample {
		fmt.Println("Key=", key, " | Value=", mapSample[key])
	}

	fmt.Println("Printing the map using key value pair")
	fmt.Println("Syntax: for key, value := range map {")
	for key, value := range mapSample {
		fmt.Printf("Key %v | Value: %v \n", key, value)
	}

	fmt.Println("")
	fmt.Println("Use Case #7: loop an array and slice using range keyword.")
	//array of firstnames
	firstNames := [4]string{"Edsel", "Jan", "Jaz", "Jil"}
	for counter, firstName := range firstNames {
		fmt.Printf("array index [%d] | value = %v \n", counter, firstName)
	}

}

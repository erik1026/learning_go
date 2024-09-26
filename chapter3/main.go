package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	// Exercise 1
	{
		fmt.Println("----- Exercise 1 -----")

		var orgSlice = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
		var subSlice1 = orgSlice[:2]
		var subSlice2 = orgSlice[1:4]
		var subSlice3 = orgSlice[3:]

		fmt.Println("Original Slice: ", orgSlice)
		fmt.Println("Subslice 1: ", subSlice1)
		fmt.Println("Subslice 2: ", subSlice2)
		fmt.Println("Subslice 3: ", subSlice3)
	}

	// Exercise 2
	{
		fmt.Println("----- Exercise 2 -----")

		message := "Hi 👩 and 👨"
		runeMsg := []rune(message)
		fmt.Println(string(runeMsg[3]))
	}

	// Exercise 3
	{
		fmt.Println("----- Exercise 3 -----")

		myStruct1 := Employee{
			"Erik",
			"Rodenburgh",
			1,
		}

		myStruct2 := Employee{
			firstName: "Abigail",
			lastName:  "Rodenburgh",
			id:        2,
		}

		var myStruct3 Employee

		myStruct3.firstName = "Bella"
		myStruct3.lastName = "Rodenburgh"
		myStruct3.id = 3

		fmt.Println(myStruct1)
		fmt.Println(myStruct2)
		fmt.Println(myStruct3)
	}

}

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Exercise 1
	{
		fmt.Println("----- Exercise 1 -----")

		intSlice := make([]int, 100)

		for idx := range intSlice {
			intSlice[idx] = rand.Intn(101)
		}

		// Book Solution
		// mySlice := make([]int, 0, 100)
		// 	for i := 0; i < 100; i++ {
		// 		mySlice = append(mySlice, rand.Intn(100))
		// 	}
		// 	fmt.Println(mySlice)

		// Exercise 2
		fmt.Println("----- Exercise 2 -----")

		for _, val := range intSlice {
			switch {
			case val%2 == 0:
				fmt.Println("Two!")
			case val%3 == 0:
				fmt.Println("Three!")
			case val%6 == 0:
				fmt.Println("Six!")
			default:
				fmt.Println("Never mind")
			}
		}

	}

	// Exercise 3
	{
		fmt.Println("----- Exercise 3 -----")
		var total int

		for i := 0; i < 10; i++ {
			total := total + i
			fmt.Println(total)
		}

		fmt.Println(total)
	}
}

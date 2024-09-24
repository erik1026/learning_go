package main

import (
	"fmt"
	"math"
)

func main() {

	// Exercise 1
	{
		var i int = 20
		var f float64 = float64(i)

		fmt.Println(i, f)
	}

	// Exercise 2
	{
		const value = 5

		var i int = value
		var f float64 = value

		fmt.Println(i, f)
	}

	// Exercise 3
	{
		var b byte = math.MaxUint8
		var smallI int32 = math.MaxInt32
		var bigI uint64 = math.MaxUint64

		fmt.Println(b + 1)
		fmt.Println(smallI + 1)
		fmt.Println(bigI + 1)
	}
}

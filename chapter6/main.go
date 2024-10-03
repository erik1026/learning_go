package main

import "fmt"

// Exercise 1
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// Exercise 2
func UpdateSlice(strSlice []string, str string) {
	fmt.Println("Inside Update Slice...")
	fmt.Println("Before update: ", strSlice)
	strSlice[len(strSlice)-1] = str
	fmt.Println("After update: ", strSlice)

}

func GrowSlice(strSlice []string, str string) {
	fmt.Println("Inside Grow Slice...")
	fmt.Println("Before Append: ", strSlice)
	strSlice = append(strSlice, str)
	fmt.Println("After Append: ", strSlice)
}

func main() {

	// Exercise 1
	{
		fmt.Println("----- Exercise 1 -----")

		person := MakePerson("Erik", "Rodenburgh", 29)
		fmt.Println(person)

		personPtr := MakePersonPointer("Erik", "Rodenburgh", 29)
		fmt.Println(*personPtr)
	}

	{
		// Exercise 2
		fmt.Println("----- Exercise 2 -----")

		strSlice1 := []string{"This", "is", "my", "slice"}
		strSlice2 := []string{"This", "is", "my", "slice"}

		UpdateSlice(strSlice1, "cat")
		fmt.Println("Main: After update function: ", strSlice1)

		GrowSlice(strSlice2, "cat")
		fmt.Println("Main: After append function: ", strSlice2)
	}

	{
		fmt.Println("----- Exercise 3 -----")

		persons := make([]Person, 0)

		for i := 0; i < 10000000; i++ {
			persons = append(persons, MakePerson("Erik", "Rodenburgh", 29))
		}
	}

}

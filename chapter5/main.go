package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func add(i int, j int) (int, error) { return i + j, nil }
func sub(i int, j int) (int, error) { return i - j, nil }
func mul(i int, j int) (int, error) { return i * j, nil }
func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}

	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func fileLen(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	buf := make([]byte, 2048)
	total := 0
	for {
		count, err := file.Read(buf)
		total += count
		if err != nil {
			if err != io.EOF {
				return -1, err
			}
			break
		}
	}

	return total, nil
}

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return prefix + " " + body
	}
}

func main() {
	// Exercise 1
	{
		expressions := [][]string{
			{"2", "+", "3"},
			{"2", "-", "3"},
			{"2", "*", "3"},
			{"2", "/", "3"},
			{"2", "/", "0"},
			{"2", "%", "3"},
			{"two", "+", "three"},
			{"5"},
		}

		for _, expression := range expressions {
			if len(expression) != 3 {
				fmt.Println("invalid expression:", expression)
				continue
			}
			p1, err := strconv.Atoi(expression[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			op := expression[1]
			opFunc, ok := opMap[op]
			if !ok {
				fmt.Println("unsupported operator:", op)
				continue
			}
			p2, err := strconv.Atoi(expression[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			result, err := opFunc(p1, p2)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(result)
		}
	}

	// Exercise 2
	{
		sizeOfFile, err := fileLen("file.txt")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Size of the file: ", sizeOfFile)
	}

	// Exercise 3
	{
		helloPrefix := prefixer("Hello")
		fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
		fmt.Println(helloPrefix("Maria")) // should print Hello Maria
	}
}

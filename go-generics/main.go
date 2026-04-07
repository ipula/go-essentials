package main

import "fmt"

func main() {
	result1 := add(3, 5)
	fmt.Println("Result of adding integers:", result1)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

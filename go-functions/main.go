package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(10, 1, 2, 3, 4, 5)
	anotherSum := sumup(20, numbers...)
	fmt.Println(anotherSum)
	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

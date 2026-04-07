package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	threeTimes := createTransformer(3)

	doubleNumbers := transformNumbers(&numbers, double)
	threeTimesNumbers := transformNumbers(&numbers, threeTimes)

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(transformed)
	fmt.Println(doubleNumbers)
	fmt.Println(threeTimesNumbers)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

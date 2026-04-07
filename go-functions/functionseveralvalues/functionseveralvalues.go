package functionseveralvalues

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}
	double := transformNumbers(&numbers, double)
	threeTimes := transformNumbers(&numbers, threeTimes)

	fmt.Println(double)
	fmt.Println(threeTimes)
	fmt.Println(numbers)

	transformedN1 := getTransfomerFn(&numbers)
	transformedN2 := getTransfomerFn(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformedN1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformedN2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func getTransfomerFn(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return threeTimes
	}
}

func transformNumbers(nums *[]int, transformFunc transformFn) []int {
	dNumbers := []int{}
	for _, val := range *nums {
		dNumbers = append(dNumbers, transformFunc(val))
	}
	return dNumbers
}

func double(num int) int {
	return num * 2
}

func threeTimes(num int) int {
	return num * 3
}

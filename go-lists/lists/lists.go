package main

import "fmt"

// func main() {
// 	prices := []float64{10.5, 20.0, 15.75, 25.0}
// 	fmt.Println("prices:", prices)

// 	fmt.Println("slice price:", prices[1:3])           // array slicing
// 	fmt.Println("Last price:", prices[len(prices)-2:]) // slicing from the end using len(prices)-2 to get the last two elements
// 	// when you slice an array, it creates a new slice that references the same underlying array. This means that changes to the slice will affect the original array and vice versa.
// 	// However, if you create a new slice that does not reference the original array, changes to that new slice will not affect the original array.
// 	featuredPrices := prices[1:3] // This creates a new slice that references the same underlying array as prices
// 	featuredPrices[0] = 30.0      // This modifies the first element of the featuredPrices slice, which also modifies the second element of the prices slice
// 	fmt.Println("prices after modification:", prices)
// 	fmt.Println("featuredPrices:", featuredPrices)
// 	fmt.Println("length of featuredPrices:", len(featuredPrices))
// 	fmt.Println("capacity of featuredPrices:", cap(featuredPrices))
// 	// The capacity of a slice is the number of elements in the underlying
// 	// array starting from the index of the first element of the slice.
// 	// In this case, since featuredPrices starts at index 1 of the prices array,
// 	// its capacity is 3 (the elements at indices 1, 2, and 3 of the prices array).
// 	// The length of featuredPrices is 2 because it contains two elements (20.0 and 15.75).
// 	fmt.Println("prices cap:", cap(prices))

// }

func main() {
	prices := []float64{10.5, 20.0, 15.75, 25.0}
	fmt.Println("prices:", prices[0:1])

	prices[0] = 12.0
	fmt.Println("prices after modification:", prices)

	updatedPrices := append(prices, 30.0)
	fmt.Println("updated prices:", updatedPrices)
	fmt.Println("original prices after append:", prices)

	discountedPrices := []float64{9.0, 18.0, 13.5, 22.5}
	prices = append(prices, discountedPrices...)
	fmt.Println("discounted prices after copy:", prices)
}

package main

import "fmt"

type floatMap map[string]float64

func (fm floatMap) output() {
	fmt.Println(fm)
}

func main() {
	userNames := make([]string, 2, 5) // Initial length of 2, but capacity of 5
	userNames[0] = "Alice"
	fmt.Println("User Names:", userNames)
	fmt.Println("Length:", len(userNames))
	fmt.Println("Capacity:", cap(userNames))

	// Adding a new element using append
	userNames = append(userNames, "Charlie")
	fmt.Println("Updated User Names:", userNames)

	// Adding more elements to exceed the initial capacity
	userNames = append(userNames, "David")
	fmt.Println("User Names after adding more:", userNames)

	course := make(floatMap, 3) // Initial capacity of 3, but it can grow dynamically
	course["Go Basics"] = 89.99
	course["Go Concurrency"] = 99.99
	course.output()
	fmt.Println("Course Map:", course)
	fmt.Println("Go Basics URL:", course["Go Basics"])

	for index, value := range userNames {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	fmt.Println("Updated Course Map:", course)
}

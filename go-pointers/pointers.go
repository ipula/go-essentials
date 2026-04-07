package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println("Age: ", age)
	fmt.Println("Age pointer: ", agePointer)           // get the address of age variable
	fmt.Println("Value at age pointer: ", *agePointer) // get the value at the pointer

	getAdultYears(agePointer)
	fmt.Println("Age: ", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}

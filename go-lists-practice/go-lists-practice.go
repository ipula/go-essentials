package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	//	1).
	hobbies := [3]string{"Reading", "Cooking", "Traveling"}
	fmt.Println("Hobbies:", hobbies)

	//2)
	fmt.Println("First hobby:", hobbies[0])
	fmt.Println("2 and 3 hobbies:", hobbies[1:])

	//3)
	mainHobbies := hobbies[0:2] // or mainHobbies := hobbies[:2]
	fmt.Println("Main hobbies:", mainHobbies)

	//4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println("Main hobbies after slicing:", mainHobbies)

	//5)\
	courseGoal := []string{"Go Basics", "Go Concurrency"}
	fmt.Println("Course Goal:", courseGoal)
	courseGoal[1] = "Go Advanced"
	fmt.Println("Updated Course Goal:", courseGoal)
	newCourseGoal := append(courseGoal, "Go Web Development")
	fmt.Println("Updated Course Goal:", newCourseGoal)

	//7)
	products := []Product{
		{id: 1, title: "Laptop", price: 999.99},
		{id: 2, title: "Smartphone", price: 499.99},
	}
	fmt.Println("Products:", products)
	newProducts := append(products, Product{id: 3, title: "Tablet", price: 299.99})
	fmt.Println("Updated Products:", newProducts)
}

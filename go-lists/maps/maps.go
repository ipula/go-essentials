package main

import "fmt"

func main() {
	// websites := []string{"https://www.google.com", "https://www.github.com"}
	// println("Websites:", websites)

	websiteMap := map[string]string{
		"Google": "https://www.google.com",
		"GitHub": "https://www.github.com",
	}
	fmt.Println("Website Map:", websiteMap)
	fmt.Println("Google URL:", websiteMap["Google"])

	// Adding a new key-value pair
	websiteMap["StackOverflow"] = "https://stackoverflow.com"
	fmt.Println("Updated Website Map:", websiteMap)

	// Deleting a key-value pair
	delete(websiteMap, "GitHub")
	fmt.Println("After deletion:", websiteMap)
}

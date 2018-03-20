package main

import "fmt"

func main() {
	var fileExtensions = map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtensions)

	delete(fileExtensions, "Kotlin")

	// delete function doesn't do anything if the key doesn't exist
	delete(fileExtensions, "Javascript")

	fmt.Println(fileExtensions)
}

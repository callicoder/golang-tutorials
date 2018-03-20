package main

import "fmt"

func main() {
	var cities = []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	var asianCities = cities[3:]
	var indianCities = asianCities[1:5]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
	fmt.Println("indianCities = ", indianCities)
}

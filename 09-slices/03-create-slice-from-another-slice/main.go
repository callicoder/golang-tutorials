package main

import "fmt"

func main() {
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	asianCities := cities[3:]
	indianCities := asianCities[1:5]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
	fmt.Println("indianCities = ", indianCities)
}

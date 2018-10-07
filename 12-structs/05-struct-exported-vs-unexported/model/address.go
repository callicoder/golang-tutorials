package model

// Unexported struct (only accessible inside package `model`)
type address struct {
	houseNo, street, city, state, country string
	zipCode                               int
}
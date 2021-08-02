package main

import (
	"flag"
	"fmt"
	"strings"
)

type hobbies []string

func (h *hobbies) String() string {
	return fmt.Sprint(*h)
}

func (h *hobbies) Set(value string) error {
	for _, hobby := range strings.Split(value, ",") {
		*h = append(*h, hobby)
	}
	return nil
}

func main() {
	// Declare a string flag called name with a default value ("Guest") and a help message
	name := flag.String("name", "Guest", "specify your name")

	// Declare a flag called age with default value of 0 and a help message
	age := flag.Int("age", 0, "specify your age")

	// Bind the command-line flag with an existing variable
	var country string
	flag.StringVar(&country, "country", "", "enter your country")

	// Defining a custom flag
	var hobbiesFlag hobbies
	flag.Var(&hobbiesFlag, "hobbies", "comma separated list of hobbies")

	// Enable command-line parsing
	flag.Parse()

	fmt.Printf("Hello %s\n", *name)
	fmt.Printf("Your age is %d\n", *age)
	fmt.Printf("Your are from %s\n", country)

	fmt.Printf("Your hobbies are: ")
	for _, hobby := range hobbiesFlag {
		fmt.Printf("%s ", hobby)
	}
	fmt.Println()

	// Read all the positional arguments
	fmt.Println("Positional Arguments:", flag.Args())

	// Read the i'th positional argument
	i := 0
	fmt.Printf("Positional argument at index %d: %v\n", i, flag.Arg(i))
}

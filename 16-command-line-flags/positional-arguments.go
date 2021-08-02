package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.Int("s", 0, "start line number")
    t := flag.Int("t", 0, "end line number")

	// Enable command-line parsing
	flag.Parse()

	fmt.Printf("Search file from line number %d to %d\n", *s, *t)

    // Read all the positional arguments
	fmt.Println("for keywords:", flag.Args())

    // Read the i'th positional argument
	i := 0
	fmt.Printf("The keyword at index %d: %v\n", i, flag.Arg(i))
}

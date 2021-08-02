package main

import (
	"fmt"
	"net/url"
)

func main() {
	// QueryUnescape: Decode a URL query string
	encodedStr := "Gol%40ng%3F%26"
	decodedStr, err := url.QueryUnescape(encodedStr)
	if err != nil {
		fmt.Printf("Error decoding the string %v", err)
	}

	fmt.Println(decodedStr)

	// PathUnescape: Decode a URL path segment
	encodedPathVar := "Gol@ng%3F&"
	decodedPathVar, err := url.PathUnescape(encodedPathVar)
	if err != nil {
		fmt.Printf("Error decoding the string %v", err)
	}
	fmt.Println(decodedPathVar)
}

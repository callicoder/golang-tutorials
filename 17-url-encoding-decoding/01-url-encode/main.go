package main

import (
	"fmt"
	"net/url"
)

func main() {
	// QueryEscape: Escape a string to safely place it inside a URL query string
	str := "Gol@ng?&"
	encodedStr := url.QueryEscape(str)

	fmt.Println(encodedStr)

	// PathEscape: Escape a string to safely place it inside a URL path segment
	pathVar := "Gol@ng?&"
	encodedPathVar := url.PathEscape(pathVar)
	fmt.Println(encodedPathVar)
}

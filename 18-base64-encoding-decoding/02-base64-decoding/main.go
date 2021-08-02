package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encodedData := "R29sQG5nIGlzIEF3ZXNvbWU/fg=="

	// Standard Base64 Decoding
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(string(decodedData))

	// URL and filename-safe Base64 decoding
	urlSafeBase64EncodedData := "R29sQG5nIGlzIEF3ZXNvbWU_fg=="
	urlSafeData, err := base64.URLEncoding.DecodeString(urlSafeBase64EncodedData)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	fmt.Println(string(urlSafeData))
}

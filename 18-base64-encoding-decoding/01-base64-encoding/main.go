package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Gol@ng is Awesome?~"

	// Standard Base64 Encoding
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedData)

	// Standard Base64 Encoding without padding
	encodedDataWithoutPadding := base64.RawStdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedDataWithoutPadding)

	// URL and filename-safe Base64 encoding
	urlSafeEncodedData := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlSafeEncodedData)

	// URL and filename-safe Base64 encoding without padding
	urlSafeEncodedDataWithoutPadding := base64.RawURLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlSafeEncodedDataWithoutPadding)

}

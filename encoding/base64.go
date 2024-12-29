package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Aziz Nur Abdul Qodir"))
	fmt.Println(encoded)

	decode, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(decode))
}

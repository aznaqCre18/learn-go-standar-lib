package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconf digunakan untuk menconvert tipe data sebuah variabel, seperti yang akan di contohkan dibawah
	/*
		strconv.parseBool(string)
		strconv.parseFloat(string)
		strconv.parseInt(string)
		strconv.FormatBool(bool)
		strconv.FormatFloat(float, … )
		strconv.FormatInt(int, … )
	*/

	// Example
	boolean, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// mereturn nilai boolean, mengecek apakah ada string di dalam string jika ada return true || false
	fmt.Println(strings.Contains("Hello World!", "Hello"))

	// mereturn slice of string, dengan isi string yg telah dipisah sesuai separator yg ditentukan, sama seperti arr.split di javascript
	fmt.Println(strings.Split("Aziz Nur Abdul Qodir", " ")[0])

	// mereturn hasil convert string menjadi lowercase
	fmt.Println(strings.ToLower("AZIZ NUR ABDUL QODIR!"))

	// mereturn hasil convert string menjadi uppercase
	fmt.Println(strings.ToUpper("aziz nur abdul qodir!"))

	// mereturn hasil convert string dengan menghilangkan apapun cutset yang ada di awal dan di akhir string
	fmt.Println(strings.Trim("     aziz nur abdul qodir!      ", " "))

	// mereturn string yang sudah di replace dengan parameter yang telah diberikan
	fmt.Println(strings.ReplaceAll("Azdev Nur Abdul", "Azdev", "Aziz"))
}

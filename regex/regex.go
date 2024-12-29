package main

import (
	"fmt"
	"regexp"
)

func main() {
	// contoh penggunaan regex

	regex := regexp.MustCompile(`a([a-z])z`)

	fmt.Println(regex.MatchString("aiz"))
	fmt.Println(regex.MatchString("auz"))
	fmt.Println(regex.MatchString("azizh"))

	fmt.Println(regex.FindAllString("aziz azz aiz auz a1z aoz atz ayz abzz acz avz auz aiz akz", 10))
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Aziz", "Umi", "Bapak", "Isti", "Hafidzh"}
	values := []int{94, 92, 89, 82, 80}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Aziz"))
	fmt.Println(slices.Index(names, "Umi"))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// func Math fungsi yang berisi sebuah kontanta atau fungsi untuk matematika

	// Example

	// ceil - membulatkan nilai float64 ke atas
	ceil := math.Ceil(162.7)
	// floor - membulatkan nilai float64 ke bawah
	floor := math.Floor(162.7)
	// round - membulatkan nilai float64 ke atas atau kebawah sesuai dengan angka mana yang paling dekat
	round := math.Round(162.7)
	// max - mengembalikan nilai float64 terbesar
	max := math.Max(32, 64)
	// min - mengembalikan nilai float64 terkecil
	min := math.Min(32, 64)

	fmt.Println(ceil)
	fmt.Println(floor)
	fmt.Println(round)
	fmt.Println(max)
	fmt.Println(min)
}

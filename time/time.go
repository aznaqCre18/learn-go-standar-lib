package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now() - untuk mendefiniskan waktu saat ini
	now := time.Now()
	fmt.Println(now)

	// time.Date(...) - untuk membuat waktu
	date := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(date)

	// time.Parse(...) -Untuk memparsing waktu dari string
	parse, _ := time.Parse(time.RFC3339, now.Format(time.RFC3339))
	fmt.Println(parse)
}

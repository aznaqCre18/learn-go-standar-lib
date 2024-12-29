package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("h", "127.0.0.1", "host")
	username := flag.String("username", "aziz", "aziznur")
	password := flag.String("password", "nur", "abdulqodir")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}

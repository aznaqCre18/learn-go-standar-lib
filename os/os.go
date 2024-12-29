package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args[0], "<<< ARGS FULL")

	for _, arg := range args {
		fmt.Println(arg, "<<< ARGS")
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname, "<<< HOSTNAME")
	} else {
		fmt.Println("Error", err.Error())
	}
}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()

	list.PushBack("Aziz")
	list.PushBack("Nur")
	list.PushBack("Abdul")
	list.PushBack("Qodir")

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := list.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

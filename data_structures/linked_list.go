package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
}

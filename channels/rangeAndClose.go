package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(cap int, c chan int) {
	x, y := 0, 1

	for i := 2; i < cap; i++ {
		n := x + y
		x = y
		y = n
		c <- n
	}
	close(c)
}

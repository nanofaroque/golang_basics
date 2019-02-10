package main

import "fmt"

func main() {
	var result = broken_calc(5, 8)
	fmt.Print(result)
}

func broken_calc(X int, Y int) int {
	count := 0
	for X < Y {
		count++
		if Y%2 == 1 {
			Y++
		} else {
			Y = Y / 2
		}
	}
	return count + X - Y
}

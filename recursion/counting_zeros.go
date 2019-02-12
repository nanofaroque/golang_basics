package main

import "fmt"

func main() {
	data := []int{1, 0, 0, 9, 8, 0,0,0,1,1,1,1}
	res:=helper(data, 0)
	fmt.Print(res)
}

func helper(data []int, pos int) int {
	if pos == len(data) {
		return 0
	}
	if data[pos] == 0 {
		return 1 + helper(data, pos+1)
	} else {
		return helper(data, pos+1)
	}
}

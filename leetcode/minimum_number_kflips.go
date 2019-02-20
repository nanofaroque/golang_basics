package main

import (
	"fmt"
)

var input []int
var input1 []int

func main() {
	input = []int{1, 1, 0}
	res := minKBitFlips(input, 2)
	fmt.Print(res)

	input1 = []int{0, 0, 0, 1, 0, 1, 1, 0}
	res1 := minKBitFlips(input1, 3)
	fmt.Print(res1)
}
func minKBitFlips(A []int, K int) int {
	pos := findFirstZeros(A)
	if pos == len(A) {
		return 0
	} // dont need any alteration
	count := 0

	for pos < len(A) {
		if A[pos] == 1 {
			pos++
		} else if A[pos] == 0 {
			if flip(pos, A, K) == -1 {
				return -1
			}
			pos++
			count++
		}
	}
	if pos < len(A) {
		return -1
	}
	return count
}
func flip(pos int, A []int, K int) int {
	if pos+K > len(A) {
		return -1
	}
	for i := pos; i < pos+K; i++ {
		if A[i] == 0 {
			A[i] = 1
		} else {
			A[i] = 0
		}
	}
	return 0

}
func findFirstZeros(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			return i
		}
	}
	return len(A)
}

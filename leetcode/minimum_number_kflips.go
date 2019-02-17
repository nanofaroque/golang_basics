package main

import "fmt"

var input []int

func main() {
	input = []int{0, 1, 0}
	res := minKBitFlips(input, 1)
	fmt.Print(res)
}
func minKBitFlips(A []int, K int) int {
	pos := findFirstZeros(A)
	count := 0
	if pos == -1 {
		return count
	}
	if pos+K > len(A) {
		return -1
	}
	//fmt.Println(pos)
	for pos >= 0 {
		pos = flip(pos, A, K)
		if pos == -1 {
			return pos
		}
		//fmt.Println(pos)
		count++
	}
	//fmt.Print(pos)
	return count
}
func flip(pos int, A []int, K int) int {
	var nextPos int
	if pos+K > len(A) {
		return -1
	}
	for i := pos; i < pos+K; i++ {
		if A[i] == 0 {
			A[i] = 1
		} else {
			A[i] = 0
			nextPos = i
		}
	}
	if nextPos == pos {
		return findFirstZeros(A)
	}
	return nextPos
}
func findFirstZeros(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			return i
		}
	}
	return len(A)
}

package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	row int
	col int
}

var matrix [][]int
var counter int
var days int

func main() {
	matrix = [][]int{
		{1,1,2,0,2,0},
	}
	counter = count(matrix)
	fmt.Print(counter)
	days = 0

	l := list.New()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 2 {
				l.PushBack(&Node{i, j})
			}
		}
	}
	fmt.Print(l.Front().Value.(*Node).row)
	fmt.Print(l.Front().Value.(*Node).col)
	res := rootten_helper(l, 0)
	fmt.Print(res)
}

func rootten_helper(l *list.List, count int) int {
	if count == counter {
		return days
	}
	if l.Len() == 0 {
		return -1
	}

	newList := list.New()
	for e := l.Front(); e != nil; e = e.Next() {
		row := e.Value.(*Node).row
		col := e.Value.(*Node).col

		if row+1 >= 0 && row+1 < len(matrix) && col >= 0 && col < len(matrix[0]) {
			if matrix[row+1][col] == 1 {
				matrix[row+1][col] = 2
				count++
				newList.PushBack(&Node{row + 1, col})
			}
		}

		if row >= 0 && row < len(matrix) && col+1 >= 0 && col+1 < len(matrix[0]) {
			if matrix[row][col+1] == 1 {
				matrix[row][col+1] = 2
				count++
				newList.PushBack(&Node{row, col + 1})
			}
		}

		if row-1 >= 0 && row-1 < len(matrix) && col >= 0 && col < len(matrix[0]) {
			if matrix[row-1][col] == 1 {
				matrix[row-1][col] = 2
				count++
				newList.PushBack(&Node{row - 1, col})
			}
		}

		if row >= 0 && row < len(matrix) && col-1 >= 0 && col-1 < len(matrix[0]) {
			if matrix[row][col-1] == 1 {
				matrix[row][col-1] = 2
				count++
				newList.PushBack(&Node{row , col - 1})
			}
		}
	}
	days += 1
	return rootten_helper(newList, count)
}

func count(data [][]int) int {
	count := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

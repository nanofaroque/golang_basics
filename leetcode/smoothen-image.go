package main

import (
	"fmt"
	"math"
)

func main() {

	var image = [][] int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
/*
	var image = [][] int{
		{1},
	}
*/
	concate(image)
}

func concate(in [][]int) [][]int {
	var row int = len(in)+2
	var col int = len(in[0])+2


	out := make([][]int, row);
	for i := range out {
		out[i] = make([]int, col)
	}
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			out[i][j] = in[i-1][j-1]
		}
	}

	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			var sum =out[i][j]+
				out[i][j-1]+out[i][j+1]+
				out[i-1][j]+out[i+1][j]+
				out[i-1][j-1]+out[i+1][j+1]+
				out[i-1][j+1]+out[i+1][j-1]
			in[i-1][j-1] = int(math.Floor(float64(sum / 9)))
		}
	}

	for i := 0; i < row-2; i++ {
		for j := 0; j < col-2; j++ {
			fmt.Print(in[i][j])
		}
		fmt.Println("")
	}

	return in
}

package main

import (
	"fmt"
	"math"
)

func main() {
	/*data:=[][]int{
		{1,1},
		{2,3},
		{5,0},
		{4,3},
	}*/

	data:=[][]int{
		{0,0},
		{1,1},
		{1,2},
	}
	fmt.Print(data)
	distance:=0
	for i:=1;i< len(data);i++  {
		distance+=minDistance(data[i-1],data[i])
	}
	fmt.Print(distance)
}

func minDistance(x []int, y []int) int {
	absX := math.Abs(float64(x[0] - y[0]))
	absY := math.Abs(float64(x[1] - y[1]))
	return int(math.Max(absX, absY))
}

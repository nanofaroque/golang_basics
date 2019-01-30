package main

import (
	"fmt"
)

func main() {
	/**
	declare and initialize an multi-dimensional array
	*/
	var in = [][] int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// printing each element
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(in[i][j])
		}
	}
	fmt.Print(in);
}

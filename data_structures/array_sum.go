package main

import "fmt"

func main()  {
	var x [5]float64

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	for i:=0;i< len(x);i++{
		fmt.Println(x[i])
	}

}

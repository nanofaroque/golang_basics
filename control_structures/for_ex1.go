package main

import (
	"fmt"
	"strconv"
)

func main(){
	i:=1 // Initialize the variable
	for i<10{
		fmt.Println(i)
		i = i+1
	}



	sum:=0;
	for j:=0;j<100;j++{
		sum+=j
		fmt.Println("I am printing the value of sum: "+strconv.Itoa(sum))
	}


}
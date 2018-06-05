package main

import "fmt"

func main()  {
	i:=10;
	switch i {

	case 0:
		fmt.Println("I am wrong");
	case 10:
		fmt.Println("I am correct ");
	default:
		fmt.Println("I am nowhere");

	}
}
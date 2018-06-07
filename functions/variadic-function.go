package main

import "fmt"
// ...    means that, add function can take multiple variale of int type
func add (args ...int) int{
	total := 0;
	for _,v :=range args{
		total +=v;
	}
	return total;
}

func main(){
	fmt.Println(add(1,2,3,4,5,6,7,8,9));
}
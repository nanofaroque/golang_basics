package main

import "fmt"
/**
You can return multiple value from a go function
*/
func f() (int, int){
	return 5,6;
}
func main(){
	x,y:=f();
	fmt.Println(x);
	fmt.Println(y);
}

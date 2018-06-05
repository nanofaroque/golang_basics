package main

import (
	"fmt"
	"time"
)
func main(){
	go goBoring("Hello");
	fmt.Println("I am listening ...");
	time.Sleep(2*time.Second);
	fmt.Println("You are boring and I am leaving ..")
}
func goBoring(s string) {
	for i:=0;i<10;i++ {
		fmt.Println(s,i);
		time.Sleep(time.Second)
	}

}
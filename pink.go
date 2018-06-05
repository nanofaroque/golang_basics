package main

import (
	"fmt"
	"time"
)

func main(){
	boring("Hello");
}
func boring(s string) {
	for i:=0;i<10;i++ {
		fmt.Println(s,i);
		time.Sleep(time.Second)
	}

}
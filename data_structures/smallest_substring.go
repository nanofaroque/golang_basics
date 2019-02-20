package main

import "fmt"

func main() {
	input:="xyyzyzyx"
	arr:=[] byte{'x','y','z'}

	fmt.Print(input)
	fmt.Print(arr)

	res:=compute(arr,input)
	fmt.Println(res)
}

func compute(arr []byte, in string) int {
	//tail:=0
	//head:=0
	m:=make(map[byte]int)

	for i:=0;i< len(arr);i++  {
		m[arr[i]]=0
	}

	fmt.Println(m)
	return 0
}

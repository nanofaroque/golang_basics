package main

import "fmt"

func main() {
	fmt.Println(circularArrayRotation([]int32{3,4,5},2,[]int32{1,2}))
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	var last int32
	for i:=0;i<(int(k));i++{
		last=a[len(a)-1]
		rightShift(a)
		a[0]=last
	}
	result:=make([]int32,len(queries))
	for i:=0;i<len(queries) ;i++  {
		pos:=queries[i]
		fmt.Println(a[pos])
		result[i]=a[pos]
	}
	return result
}

func rightShift(arr []int32) {
	l:=len(arr)
	copy(arr[1:l], arr[0:l])
}

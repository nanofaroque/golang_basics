package main

import "log"

func main() {
	result:=compareTriplets([]int32{5,6,7},[]int32{3,6,10})
	log.Println(result[0],result[1])
}
func compareTriplets(a []int32, b []int32) []int32 {
	triplets:=make([]int32 ,2)
	alice := 0
	alice=0
	bob:=0
	for i:=0;i<3;i++{
		if a[i]>b[i]{
			alice+=1
		}else if b[i]>a[i]{
			bob+=1
		}
	}
	triplets[0]= int32(alice)
	triplets[1]= int32(bob)
	return triplets
}

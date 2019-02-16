package main

import (
	"testing"
)

type CompareTest struct {
	x []int32
	y []int32
	expected []int32
}

var cases= CompareTest{
	[]int32{5,6,7},[]int32{3,6,10},[]int32{1,1},
}

func Test_compareTriplets(t *testing.T)  {
	result:=compareTriplets(cases.x,cases.y)
	if result[0]!=cases.expected[0]{
		t.Fatal("fatal error")
	}

	if result[1]!=cases.expected[1]{
		t.Fatal("fatal error")
	}

}

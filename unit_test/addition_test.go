package main

import "testing"

type AdditionTest struct {
	x int
	y int
	expected int
}

var cases=[] AdditionTest{
	{2,2,4},
	{2,1,3},

}

func Test_Add(t *testing.T)  {
	for _,test:=range cases{
		result:=add(test.x,test.y)
		if result!=test.expected{
			t.Fatal("fatal error")
		}
	}
}


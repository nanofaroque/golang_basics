package main

import (
	"fmt"
	"testing"
)

func main()  {
	diStringMatch("IDID")
	diStringMatch("III")
	diStringMatch("DDI")
}

func diStringMatch(S string) []int {
	var result = make([]int, len(S)+1)
    var nextMax=len(S)
    var nextMin=0
    for i:=0;i<len(S);i++{
    	if S[i]=='D'{
    		result[i]=nextMax
    		nextMax--
		}
		if S[i]=='I'{
			result[i]=nextMin
			nextMin++
		}
	}
    result[len(S)]=nextMax
    fmt.Print(result)
	return result
}

func di_string_match_test(t *testing.T)  {
	var output = diStringMatch("IDID")
	if output[0] != 0 {
		t.Error("Expected 1.5, got ", output[0])
	}
}


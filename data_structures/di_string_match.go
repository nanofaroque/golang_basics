package main

import "fmt"

func main()  {
	diStringMatch("IDID")
	diStringMatch("III")
	diStringMatch("DDI")
}

func diStringMatch(S string) []int {
	var result = make([]int, len(S)+1)
    var nextMax=len(S)
    var nextMin=0
    //var i int
    for i:=0;i<len(S);i++{
    	if S[i]=='D'{
    		result[i]=nextMax
    		nextMax--
		}

		if S[i]=='I'{
			result[i]=nextMin
			nextMin++
		}
		//fmt.Print(result)
	}
    result[len(S)]=nextMax
    fmt.Print(result)
	return result
}

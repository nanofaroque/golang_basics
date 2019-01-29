package main

import "fmt"

func main() {
	// declare
	var m map[string]int
	// initialize map
	m = make(map[string]int)
	// input string
	var s = "hello world"
	// iterate and count
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]) + "->")
		if m[string(s[i])] == 0 {
			m[string(s[i])] = 1
		} else {
			m[string(s[i])] = m[string(s[i])] + 1
		}
	}
	//print map
	fmt.Println(m)
}

package main

import "strings"

func main() {
	str := "leet2code3"
	var sb strings.Builder
	decodeStr(str, 0, sb)

}

func decodeStr(in string, pos int, ret strings.Builder) string {
	if pos == len(in) {
		return ret.String()
	}

}

package main

import "fmt"

func main() {
	data := [] int{100, 4, 200, 3, 2, 1}
	fmt.Print(data)
	m := make(map[int]bool)

	for i := 0; i < len(data); i++ {
		m[data[i]] = true
	}
	fmt.Println(m[308])
	max:=0
	for i := 0; i < len(data); i++ {
		if m[data[i]-1]==false{
			cur:=0
			target:=data[i]
			for m[target]==true  {
				target=target+1
				cur=cur+1
			}
			if cur>max{
				max=cur
			}
		}
	}

	fmt.Println("Max",max)

}

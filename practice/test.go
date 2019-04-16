package main

import (
	"fmt"
	"golang_basics/practice/interfaces"
)
func main() {
	fmt.Print("hello")
	animals := []interfaces.Animal{interfaces.Dog{}, interfaces.Cat{}}
	for i, animal := range animals {
		fmt.Println(i)
		fmt.Println(animal.Speak())

	}
}

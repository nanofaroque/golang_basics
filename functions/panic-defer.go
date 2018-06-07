package main

import "fmt"

func main() {
	defer func() {
		int := recover()
		fmt.Println(int)
	}()
	panic(100)
}

//Earlier we created a function that called the panic function to cause a run time error. We can handle a run-time
// panic with the built-in recover function. recover stops the panic and returns the value that was passed to the call to panic.
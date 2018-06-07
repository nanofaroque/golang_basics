package main
import "fmt"

func first(){
	fmt.Println("I am from first")
}

func second()  {
	fmt.Println("I am from second")
}

//Go has a special statement called defer which schedules a function call to be run after the function completes
func main(){
	defer second()
	first()
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}
//This has 3 advantages: (1) it keeps our Close call near our Open call so it's easier to understand,
// (2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them and
// (3) deferred functions are run even if a run-time panic occurs.
package main

import "time"

func main()  {
	for i:=0; i<24; i++{
		c:= timer(1*time.Second);
		<-c; // writing to the channel

	}
}

func timer(d time.Duration) <-chan int{
	c := make(chan int) // since you will receive and send data through c, you have to initialize as a channel
    go func() {
    	time.Sleep(d)
    	c<- 1
	}()

	return c
}

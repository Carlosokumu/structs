package main

/*
     Channels provide us with a way to pass Data between different goroutines since they
	 run concurrently and we won't know when one is finished
*/
import (
	"fmt"
)

func main() {

	// This is where we "make" the channel, which can be used
	// to move the `int` datatype
	out := make(chan int)

	go multiplyByTwo(6, out)
	/*


	   // We pause the program so that the `multiplyByTwo` goroutine
	   // can finish and print the output before the code exits
	   //time.Sleep(time.Second)


	*/
	// Once any output is received on this channel, print it to the console and proceed
	fmt.Println(<-out)
}

//Lets create a function to multiply a nuumber by two then return its result
//The function now accepts a channel of type Int
func multiplyByTwo(n int, out chan<- int) int {
	result := n * 2
	//Lets put result into the channel
	out <- result
	return result
}

/*
Soham Kamani
 Directional Channels ->You restrict a channel that it is only available to send data or only
 receive data
*/

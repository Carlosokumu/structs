package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {

	//go hello()
	//time.Sleep(1 * time.Second)
	//fmt.Println("main function")
	start := time.Now()
	//Self executing function
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()
	//self executing function
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println(start)
	elapsed := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsed.String())
	time.Sleep(time.Minute)

}

package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}

}
func multiplyByTwo(n int) int {
	result := n * 2
	println(result)
	return result
}
func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	/*
	   A goroutine is a function capable of running concurrently with other functions
	*/
	go f(4)
	var input string
	fmt.Scanln(&input)
	/*
	   Channels provide a way for two goroutines to communicate
	*/
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input2 string
	fmt.Scanln(&input2)

}

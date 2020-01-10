package main

import "fmt"

func main() {
	c := make(chan string)
	fmt.Println("Sending to the channel")
	go func(input chan string) {
		input <- "Hello"
	}(c)

	fmt.Println("recieving from the channel")
	greeting := <-c
	fmt.Println("greeting channel")
	fmt.Println(greeting)

}

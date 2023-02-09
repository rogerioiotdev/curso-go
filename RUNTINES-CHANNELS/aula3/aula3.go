package main

import "fmt"

//thread 1

func main() {

	//thread 1 <-> 2

	hello := make(chan string)

	//thread 2

	go func() {
		hello <- "Hello World"
	}()

	result := <-hello
	fmt.Println(result)

}

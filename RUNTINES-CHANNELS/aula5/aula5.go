package main

import (
	"fmt"
	"time"
)

func main() {

	hello := make(chan string)

	go func() {

		hello <- "hello world"

	}()

	select {

	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("default")
	}

	time.Sleep(time.Second * 5)
}

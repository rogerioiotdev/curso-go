package main

import (
	"fmt"
	"time"
)

func main() {

	queve := make(chan int)

	go func() {

		i := 0
		for {
			time.Sleep(time.Second)
			queve <- i
			i++
		}
	}()

	for x := range queve {

		fmt.Println(x)
	}
}

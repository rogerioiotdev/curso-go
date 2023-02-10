package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)

		//fmt.Println("Fim")
	}
}

func main() {
	go contador("sem go routine")
	go contador("com go routine")
	time.Sleep(time.Second * 10)

	fmt.Println("Fim")

}

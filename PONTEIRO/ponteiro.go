package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func main() {
	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)

	carro := Carro{
		Name: "Ka",
	}

	carro.andou()
	fmt.Println(carro.Name)
}
func abc(a *int) {

	*a = 200
}

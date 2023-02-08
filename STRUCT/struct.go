package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)

}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func main() {

	cliente := Cliente{
		Nome:  "Rogerio",
		Email: "rogerio@rogerio.com",
		CPF:   15234813800,
	}
	fmt.Println(cliente)
	cliente2 := Cliente{"Karina", "karina@karina.com", 2153240600}

	fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Fabio",
			Email: "fabio@frb.com",
			CPF:   1234567800,
		},

		Pais: "Fr",
	}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d. Pais: %s\n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	jsoncliente4 := `{"Nome":"Fabio","Email":"fabio@frb.com","CPF":1234567800,"pais":"Fr"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsoncliente4), &cliente4)
	fmt.Println(cliente4)
}

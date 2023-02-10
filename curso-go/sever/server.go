package main

// Imports necessários
import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Tomada de ações ao receber requisições do ESP
func handleConnection(c net.Conn) {

	// Variável booleana de estado do relé
	state := true

	// Exibe mensagem informando qual IP está sendo utilizado pelo host
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	// Enquanto não houver erros
	for {

		// Verifica novas mensagens
		netData, err := bufio.NewReader(c).ReadString('\n')

		// Se ocorreu algum erro, aborta função
		if err != nil {
			fmt.Println(err)
			return
		}

		// Se recebeu a mensagem "STOP", aborta conexão com o client correspondente
		temp := strings.TrimSpace(string(netData))

		if temp == "STOP" {
			break
		}

		// Envia o estado para o cliente correspondente
		if state {
			c.Write([]byte(string("ON")))
		} else {
			c.Write([]byte(string("OFF")))
		}

		// Muda de estado
		state = !state
	}
	// Fecha a conexão
	c.Close()
}

func main() {
	// Porta de conexão
	PORT := ":8001"

	// Inicia o listen
	l, err := net.Listen("tcp4", PORT)
	// Se acontecer algum erro, aborta o programa
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listen ok")

	// Aguarda conexões (Accept)

	defer l.Close()
	fmt.Println("Waiting for connections...")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Para uma nova conexão, executa os comandos através da função handleConnection
		go handleConnection(c)
	}
}

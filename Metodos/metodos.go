package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func main() {
	usuario1 := usuario{"João Pedxx", 20}
	usuario1.salvar()

}

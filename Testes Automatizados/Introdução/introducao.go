package main

import (
	"fmt"

	enredecos "introducao.go/enderecos"
)

func main() {
	tipoEndereco := enredecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)

}

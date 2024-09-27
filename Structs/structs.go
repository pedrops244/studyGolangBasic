package main

import "fmt"

type user struct {
	nome  string
	idade int
}

func main() {
	fmt.Println("Arquivo structs")

	// var u user
	// u.nome = "nome"
	// u.idade = 21

	// fmt.Println(u)

	usuario2 := user{"Pedro", 22}
	fmt.Println(usuario2)

}

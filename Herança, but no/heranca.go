package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	e1 := estudante{pessoa{"Pedro", "ksks", 22, 178}, "Engenharia", "ISP"}

	fmt.Println(e1)
}

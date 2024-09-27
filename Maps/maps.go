package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])
}

package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando..")
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)

	// for j := 0; j <= 10; j++ {
	// 	fmt.Println("Incrementando", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Luas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}

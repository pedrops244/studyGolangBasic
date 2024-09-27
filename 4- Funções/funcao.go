package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println("Função f")
		return text
	}
	result := f("Texto da função 1")
	fmt.Println(result)

	rm, _ := calculosMatematicos(10, 15)
	fmt.Println(rm)
}

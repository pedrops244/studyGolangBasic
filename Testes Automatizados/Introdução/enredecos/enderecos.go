package enredecos

import (
	"fmt"
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return fmt.Sprintln(strings.Title(primeiraPalavraDoEndereco), "é um tipo válido")
	}

	return "Tipo inválido"
}

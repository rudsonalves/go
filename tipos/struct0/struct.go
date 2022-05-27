package main

import "fmt"

type pessFisica struct {
	nome       string
	nascimento string
	cpf        int64
}

type pessJuridica struct {
	cnpj         int64
	fantasia     string
	endComercial string
	pessFisica
}

func main() {
	p1 := pessFisica{
		nome:       "Alberto Santos",
		nascimento: "12/05/1978",
		cpf:        87344698552,
	}

	p2 := pessFisica{
		nome:       "Elisangela Sampaio",
		nascimento: "27/01/1985",
		cpf:        95478523344,
	}

	fmt.Println(p1, p2)

	emp := pessJuridica{
		cnpj:         35879964000189,
		fantasia:     "Artigos Esportivos Santos",
		endComercial: "Rua Saturnino Rui, 125, Santos, SP",
		pessFisica:   p2,
	}

	fmt.Println(emp)

}

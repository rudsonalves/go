package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func lerArquivo(nomeArq string) (texto string) {
	file, err := os.Open(nomeArq)
	if err != nil {
		log.Fatalf("falha na abertura do arquivo %s", nomeArq)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line_bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		texto += string(line_bytes) + "\n"
	}
}

func main() {
	var (
		count5  = [3]int{}
		remover = []string{"\n", "\t", ",", ".", "—"}
	)

	const arq = "albert.txt"

	// ler o texto e remove trima seu conteúdo
	texto := strings.TrimSpace(lerArquivo(arq))

	fmt.Println(texto)

	for _, c := range remover {
		texto = strings.Replace(texto, c, " ", -1)
	}

	// remove 2 ou mais espaços consecutivos
	lenText := 0
	for lenText != len(texto) {
		lenText = len(texto)
		texto = strings.Replace(texto, "  ", " ", -1)
	}

	// divide em uma slice por palavras
	palavras := strings.Split(texto, " ")

	for _, palavra := range palavras {
		switch comprimento := len(palavra); comprimento {
		case 0:
			continue
		case 1, 2, 3, 4, 5:
			count5[0]++
		case 6, 7, 8, 9, 10:
			count5[1]++
		default:
			count5[2]++
		}
	}

	for _, palavra := range palavras {
		switch comprimento := len(palavra); {
		case comprimento == 0:
			continue
		case comprimento <= 5:
			count5[0]++
		case comprimento <= 10:
			count5[1]++
		default:
			count5[2]++
		}
	}

	fmt.Printf("\n%2d palavras com até 5 letras\n", count5[0])
	fmt.Printf("%2d palavras com 6 a 10 letras\n", count5[1])
	fmt.Printf("%2d palavras com mais de 10 letras\n", count5[2])
}

package matematica

import (
	"fmt"
	"strconv"
)

// Média retorna a média dos valores passados
func Media(numeros ...float64) float64 {
	soma := 0.0
	for _, num := range numeros {
		soma += num
	}
	media := soma / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}

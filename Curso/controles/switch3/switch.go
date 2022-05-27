package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func imprimeTipo(i interface{}) {
	fmt.Println("É tipo: " + tipo(i))
}

func main() {
	imprimeTipo(2)
	imprimeTipo(2.67)
	imprimeTipo("Dois")
	imprimeTipo(time.Now())
	imprimeTipo(func() {})
}

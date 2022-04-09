package main

import "fmt"

func init() {
	fmt.Println("Inicializando 2...")
}

/* Execute via terminal com o comando

$ go run *.go

As duas funções init() serão executadas antes da função
main() no primeiro init.go

$ go run *.go
Inicializando 2...
Inicializando 1...
Main...
$
*/

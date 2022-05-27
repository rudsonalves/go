package main

import (
	"fmt"
	"time"
)

type Contador struct {
	total             int
	ultimaAtualizacao time.Time
}

func (c *Contador) Incrementa() {
	c.total++
	c.ultimaAtualizacao = time.Now()
}

func (c Contador) String() string {
	return fmt.Sprintf("total: %d, última atualização: %v", c.total, c.ultimaAtualizacao)
}

func main() {
	var c Contador
	fmt.Println(c.String())
	c.Incrementa()
	fmt.Println(c.String())
}

// Estruturas são coleções de "campos"
//Serve para agrupar dados e formar registros
//Forma da estrutura -> "type (nome da estrutura) struct"

package main

import "fmt"

type pessoa struct {
	nome   string
	idade  int
	altura float64
	genero string
}

func main() {

	fmt.Println(pessoa{"Alexandre", 24, 1.75, "Masculino"})
	fmt.Println(pessoa{"Isabela", 21, 1.70, "Feminino"})
	fmt.Println(pessoa{"Muringa", 50, 1.75, "Masculino"})
}

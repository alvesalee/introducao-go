package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

//Este método area possui um tipo 'retangulo'
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 50, altura: 25}

	//Aqui chamamos os 2 metodos definidos para nossa estrutura

	fmt.Println("Area:  ", r.area())
	fmt.Println("Perimetro:  ", r.perimetro())
}

//Metodo é uma função associada a um tipo particular
//pega uma variavel e coloca uma funcao para ela, um valor especifico

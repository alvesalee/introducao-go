package main

import "fmt"

var a int = 50
var b string = "Nome"

func main() {
	//fazer a conversão de tipo

	var c float64 = float64(a)
	fmt.Println("O valor de 'C' é:", c, "e o valor de 'B' é:   ", b)

}

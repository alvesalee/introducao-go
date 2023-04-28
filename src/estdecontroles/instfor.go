package main

import "fmt"

func main() {
	//fmt.Println("1")
	//fmt.Println("2")
	//fmt.Println("3")
	//fmt.Println("4")
	//fmt.Println("5")
	//fmt.Println("6")
	//fmt.Println("7")
	//fmt.Println("8")
	//fmt.Println("9")
	//fmt.Println("10")

	//Utilizando crase `` ele imprime do jeito que foi digitado, exemplo:

	// fmt.Println(`1,2,3,4,5,6,7,8,9,10`)

	// Ou então ...

	//fmt.Println(`1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10`)

	//Começando a utilizar o for

	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i = i + 1
	//
	//	}

	// Outra forma de escrever utilizando o for

	for i := 1; i <= 25; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	//comando IF "SE" vai falar se o numero é impar ou par
	// resto em linguagem de programação é representado pelo sinal de porcentagem " % "
	// Exemplo if
	//
	//se i%2 ==0
	//imprimir: numero par
	//se nao
	// imprimir : numero impar
}

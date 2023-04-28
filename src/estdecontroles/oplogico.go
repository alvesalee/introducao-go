package main

import "fmt"

func main() {
	//	x := 2
	//	if x == 2 || x == 3 {
	//		fmt.Println("Sim, X é igual a 2 OU 3.")
	//
	//	} else {
	//		fmt.Println("Negativo, não é igual 2 nem 3.")
	//	}

	x := 6
	if x%2 == 0 && x%3 == 0 {
		fmt.Println("Sim, x = ", (x), "é multiplo de 2 e 3!")

	} else {
		fmt.Println("X Não é multiplo de 2 e 3")
	}

}

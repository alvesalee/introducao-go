// Jeitos de declarar fatia em go
// var x [] float64
// fatia:=make ([]float64, 4)
// fatia:=[low:high]
// fatia:= arr[0:5]

package main

import "fmt"

func main() {

	// arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	// fatia := arr[2:6] // fatia é o que está por dentro da array
	// fmt.Println(fatia)

	// fatia1 := []int{1, 2, 3}
	// fatia2 := append(fatia1, 2, 4, 5)
	// fmt.Println(fatia1, fatia2)
	// Irá mostrar o antigo arranjo e o novo com os novos elementos em sequencia

	//Usando comando COPY
	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 2) // vai colocar o numero de quantos elementos é para ser copiado
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}

// Se eu quiser criar/acrescentar algum elemento usar o comando APPEND
// copiar uma fatia para outra, usar o comando COPY

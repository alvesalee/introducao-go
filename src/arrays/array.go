package main

import (
	"fmt"
)

func main() {
	var x [5]float64
	x[0] = 10
	x[1] = 5
	x[2] = 2.5
	x[3] = 6
	x[4] = 4

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]

	}
	fmt.Println(total / float64(len(x)))

}

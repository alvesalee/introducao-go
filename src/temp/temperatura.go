package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Printf("A temperatura da água em °F = %g . A temperatura da água em °C = %g ", tempF, tempC)

}

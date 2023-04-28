package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Horas:  ", horas)

		for minutos := 0; minutos < 60; minutos++ {

			fmt.Println("Minutos:  ", minutos)
		}
	}
}

package main

// Desafio de código para fazer conversões de escala termométricas.
import "fmt"

//Ponto de ebulição da agua em cada escala:
//Celsius = 100°C
//Kelvin = 373 K
//Fahrenheit = 212 F

const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK
	tempC := tempK - 273.0

	fmt.Printf("A temperatura de ebulição da agua em Kelvin é: %g , e a temperatura da água em graus Celsius é : %g ", tempK, tempC)
}

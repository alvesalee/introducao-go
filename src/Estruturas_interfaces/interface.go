package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
	//Aqui temos uma interface para formas geometricas
}

// area do quaddrado é lado² ou lado*lado
//perimetro é a soma dos lados :D

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
	//area do circulo é (phi - 3,14...)*raio² perimetro do circulo : 2*r*(phi)
}

// Para usar interfaces em go é preciso usar todos os metodos da interface. Aqui nós usaremos 'geometria' em 'quadrado'
func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// Agora a implementação do circulo
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

/* Se a variavel tem um tipo interface,  então nós podemos chamar metodos que estao na interface nomeada.
Aqui temos uma função generica 'medida' tomando vantagem desse trabalho em qualquer 'geometria'
*/

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {

	q := quadrado{lado: 10}
	c := circulo{raio: 5}

	/* os tipos de estrutura 'circulo e 'quadrado' ambos implementam a interface 'geometria'
	então nós podemos usar instancias dessas estruturas como argumentos para medir */

	medir(q)
	medir(c)
}

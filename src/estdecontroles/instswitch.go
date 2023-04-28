package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		//if i == 0 {
		//	fmt.Println("Zero")
		//} else if i == 1 {
		//	fmt.Println("Um")
		//} else if i == 2 {
		//	fmt.Println("Dois")
		//} else if i == 3 {
		//	fmt.Println("Três")
		//} else if i == 1 {
		//	fmt.Println("Um")
		//}

		// Agora fazendo com switch, uma formai mais organizada e simples

		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
		case 4:
			fmt.Println("Quatro")
		case 5:
			fmt.Println("Cinco")
		}
	}

}

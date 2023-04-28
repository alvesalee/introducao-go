package main

import "fmt"

func main() {

	/*	x := make(map[string]int)
		x["chave"] = 10
		fmt.Println(x["chave"]) */

	/*	x := make(map[int]int)
		x[1] = 30
		x[2] = 30
		fmt.Println(x[1], x[2])
	*/

	serie := make(map[string]string)
	serie["br"] = "breaking"
	serie["bd"] = "bad"
	serie["vk"] = "vikings"
	serie["bruxo"] = "The witcher"

	fmt.Println(serie["br"], serie["bd"], serie["vk"], serie["bruxo"])
}

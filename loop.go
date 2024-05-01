package main

import (
	"fmt"
)

func main() {
	var nomes []string
	nomes = append(nomes, "Ronaldo")
	nomes = append(nomes, "Olivia")
	nomes = append(nomes, "Diego")
	nomes = append(nomes, "Camila")

	for index, nome := range nomes {
		fmt.Println(index, nome)
	}

	multiplicador := 8

	fmt.Println("\n Tabuada do ", multiplicador)
	for i := 0; i < 101; i++ {
		fmt.Println(multiplicador, " X ", i, " = ", multiplicador*i)

	}

}

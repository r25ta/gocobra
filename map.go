package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["route"] = 66

	i := m["route"]

	x, y := m["route"]

	fmt.Println("Valor de map m:", m)
	fmt.Println("Valor de map i:", i)
	fmt.Println("Valor de x: ", x)
	fmt.Println("Valor de y:", y)

	//inicialize with any values

	porsches := map[string]int{
		"cayman  ": 275,
		"carrera ": 293,
		"gt3     ": 296,
		"dakar   ": 240,
	}
	println("\n Lista de porsches ")
	println("\n Modelo   | Velocidade")
	println("========================")
	for i, porsche := range porsches {
		println(i, " | ", porsche)
	}
}

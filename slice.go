package main

import (
	"fmt"
)

// SLICE DON'T NEED TO DECLARE THE ARRAY SIZE
func main() {
	sliceint := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice Int: ", sliceint)

	sliceString := []string{"Ronaldo", "Olivia", "Diego", "Camila"}
	fmt.Println("Slice String: ", sliceString)
	//USING APPEND TO ADD SLICE ELEMENTS
	sliceString = append(sliceString, "Artur")
	sliceString = append(sliceString, "Telma")

	fmt.Println("Slice Sting with append two elements: ", sliceString)

	slicefloat := []float64{11.2, 3340.9, 12345.987, 67849.9877}
	fmt.Println("Slice Float: ", slicefloat)
	fmt.Printf("Slice Float Format (#.###): %.3f ", slicefloat)
}

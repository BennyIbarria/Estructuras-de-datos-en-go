package main

import (
	"fmt"
	"os"
)

func readArray() []int {
	var n int
	fmt.Fprint(os.Stdout, "Capture el numero de elementos que contendra´el arreglo: ")
	fmt.Fscan(os.Stdin, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fprint(os.Stdout, "Capture el elemento numero ", i+1, " Del arreglo: ")
		fmt.Fscan(os.Stdin, &arr[i])
	}
	return arr
}

func sum1dArray(array []int) []int {
	for i := 1; i < len(array); i++ {
		array[i] += array[i-1]
	}
	return array
}
func main() {
	test := readArray()
	fmt.Println("Numeros capturados por el usuario: ", test)
	sum1dArray(test)
	fmt.Println("Nuevo Arreglo: ", test)
}

package main

import "fmt"

func main(){
	 arreglo := [3]int{1,2}

	 arreglo[2] = 20

	for i := 0; i < len(arreglo); i++  {
		fmt.Println(arreglo[i])
	}

	var matriz [3][2]int

	fmt.Println(matriz)
}

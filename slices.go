package main

import (
	"fmt"
)

func main()  {
	//un slice no tiene limite definido
	matriz := []int{1,2,3,4,5}
	//matriz_vacia := []int
	fmt.Println(matriz)


	/*if matriz_vacia == nil{
		fmt.Println("Esta vacio")
	}else {
		fmt.Println("no  esta vacio")
	}*/
	//puntero al arreglo
	//longitud del arreglo al que apunta
	//capacidad
	arreglo := [3]int{1,2,3}
	slice := arreglo[:2]//esta forma de mostrar el slice  parte el arreglo basado en los datos que se le pasan
	fmt.Println(slice)

}

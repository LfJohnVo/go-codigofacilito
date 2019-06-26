package main

import (
	"fmt"
)

func main()  {
	//make: tipo de slice, despues longitud, capacidad (opcional)
	slice := []int{1,2,3,4}
	copia := make([]int,4)
	//insertar informacion
	//slice = append(slice,2)
	//cap devuelve la capacidad del slice
	copy(copia,slice)

	fmt.Println(slice)
	fmt.Println(copia)


}


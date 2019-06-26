package main

import "fmt"

func main()  {
	/*
	1. Es una direccion de memoria
	2. En lugar del valor, tenemos la direccion en la que esta el valor
	3. x,y => as123d => 5
	4. x =>  as123d => 6
	5. Y ¿? => 6
	Define variable
	*T => *int, *string, *Struct
	Valor zero => nil
	*/

	var x,y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(y)

}

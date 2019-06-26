package main

import "fmt"

func main(){
	for i:=0; i<10; i++  {
		fmt.Println("Hola mundo : ", i)
	}

	//bucle infinito para remplazar al while
	/*i:=0
	for{
		fmt.Println(i)
		i++
		//Con esto se rompe el ciclo al llegar a 10
		if i > 10{
		break
	}
	}*/
}

package main

import (
	"fmt"
	"os"
	//Pagueteria que manejar buffer de memoria automaticamente
	"bufio"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre,err := reader.ReadString('\n')

	if(err != nil){
		fmt.Println(err)
	}else {
		fmt.Println("Hola " + nombre)
	}

	//Con esta tomo valores de la consola
	/*var edadr int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n",&edadr)
	fmt.Printf("Mi edad es %d\n", edadr)
	var nombre string
	fmt.Println("Coloca tu nombre: ")
	fmt.Scanf("%v\n",&nombre)
	fmt.Printf("Mi nombre es %s\n", nombre)
*/
/*
	edad := 22
	nombre := "John"
	fmt.Printf("Mi edad es %v ", edad)
	fmt.Printf(" Y mi nombre es: %v ", nombre)
*/

}

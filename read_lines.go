package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	executeReadFile()
	fmt.Println("Nunca me voy a imprimir")
}

func executeReadFile(){
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, error := os.Open("./holaa.txt")
	defer func() {
		archivo.Close()
		fmt.Println("Defer")
		//Con esto indica el error con el cual se paniqueo el sistema
		r := recover()
		fmt.Println(r)

	}()

	if error != nil{
		panic(error)
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan(){
		i++
		linea := scanner.Text()
		fmt.Println(linea)
		fmt.Println(i)

	}
	if true {
		return true
	}

	fmt.Println("nunca me ejecuto")
	archivo.Close()
	return true
}
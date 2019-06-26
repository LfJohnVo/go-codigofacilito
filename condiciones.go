package main

import "fmt"

func main(){
	x := 10
	y := 10

	if x > y {
		fmt.Printf("%d es mayor que %d \n",x,y)
	}else if x < y {
		fmt.Println("Entre al else if")
	} else{
		fmt.Println("son iguales")
	}
}//temrina func
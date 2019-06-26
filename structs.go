package main

import "fmt"

type User struct {
	edad int
	nombre , apellido string

}

func (this User) nombre_completo() string{
	return this.nombre + " " + this.apellido
}

func (this *User) set_name(n string){
	this.nombre = n
}

func main()  {
	//var uriel User
	//fmt.Println(uriel)
	//otra manera de hacerlo
	//fmt.Println(User{nombre: "Uriel", apellido:"Hernandez"})

	uriel := new(User)
	uriel.nombre = "Uriel"

	uriel.set_name("Marcos")

	uriel.apellido = "Hernandez"


	fmt.Println(uriel.nombre_completo())
	fmt.Println(uriel.nombre)

}

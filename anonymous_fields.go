package main

import "fmt"

type Human struct{
	name string
}

type Dummy struct{
	name string
}

type Tutor struct{
	Human
	Dummy
}

func (this Human) hablar() string {
	return "bla bla bla "
}

func (this Tutor) hablar() string {
	return this.Human.hablar() + "Bienvenidos a codigo facilito"
}

func main(){
	tutor := Tutor{Human{"Uriel"} , Dummy{"Uriel"}}

	//Esto sirve para especificar el atributo y la estructura que se va a usar
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.hablar())
}
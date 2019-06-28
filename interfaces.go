package main

import "fmt"

type User interface {
	Permisos() int //nivel de permiso 1 al 5
	Nombre() string
}

type Admin struct {
	name string
}

func (this Admin) Permisos() int  {
	return 5	
}

func (this Admin) Nombre() string{
	return this.name
}

func  auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + "tiene permisos de administrador"
	}else if permisos == 3{
		return user.Nombre() + "tiene permisos de editor"
	}
	return ""
}

type Editor struct {
	name string
}

func (this Editor) Permisos() int  {
	return 3
}

func (this Editor) Nombre() string{
	return this.name
}


func main() {
	//admin := Admin{"Uriel "}
	//editor := Editor{"Fulano "}
	usuarios := []User{Admin{"Uriel "},Editor{"Fulano "}}

	for _,usuario := range usuarios{
		fmt.Println(auth(usuario))
	}

}
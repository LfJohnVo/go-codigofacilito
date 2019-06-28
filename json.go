package main

import (
	"encoding/json"
	"net/http"
)

//Los atributos siempre deben empezar con mayuscula
type Curso struct{
	Title string 		`json:"title"`
	NumeroVideos int	`json:"numero_videos"`
}
//este es un arreglo de curso y se llamara cursos
type Cursos []Curso

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {


		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Ruby", 20},
			Curso{"Curso de NodeJS", 40},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8080" , nil)

}
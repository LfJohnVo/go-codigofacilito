package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/holamundo" , holamundo)
	http.HandleFunc("/" , handler)
	http.ListenAndServe(":8080",nil)
}

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("Hay una nueva peticion")
	io.WriteString(w, "Hola mundo")
}

func holamundo(w http.ResponseWriter,r *http.Request){
	fmt.Println("Hay una nueva peticion en hola mundo")
	io.WriteString(w, "HOOOLAAAAAAA")
}
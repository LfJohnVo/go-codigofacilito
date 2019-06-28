package main

import (
	"net/http"
)

func main() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.Path)
		//http.ServeFile("index.html")
		//http.ServeFile(w,r,r.URL.Path[1:])
	})*/

	//Servir archivos dentro de public
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/",fileServer))

	http.ListenAndServe(":8080" , nil)
}
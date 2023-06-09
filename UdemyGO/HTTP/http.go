package main

import (
	"log"
	"net/http"
)

func main() {
	//HTTP é um protocolo de comunicação  - Base da comunicação WEB
	// CLIENTE - SERVIDOR

	// Request - Response
	/*Rotas
	  URI -Idantificador do Recurso
	  Metodo - GET, POST, PUT, DELETE
	*/
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) { //("URI",Func())
		w.Write([]byte("Ola Mundo!"))
	})
	log.Fatal(http.ListenAndServe(":5000", nil))

}

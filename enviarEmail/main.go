package main

import (
	"fmt"
	"net/http"

	"github.com/brunobog/enviarEmail/manipulador"
)

func main() {
	fmt.Println("O serviço subiu")

	http.HandleFunc("/", manipulador.Ola)
	http.HandleFunc("/mail", manipulador.EnviaEmail)

	http.ListenAndServe(":8080", nil)
	fmt.Println("O serviço subiu")

}

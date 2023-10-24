package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Seja bem vindo a minha HOME")

}
func Contatos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Seja bem vindo a minha pagina de CONTATO")

}

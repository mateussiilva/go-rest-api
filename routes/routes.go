package routes

import (
	"log"
	"net/http"

	"github.com/mateussiilva/go-rest-api/controllers"
)


func HandleRequest(){
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/contatos",controllers.Contatos)
	log.Fatal(http.ListenAndServe(":9000",nil))
}


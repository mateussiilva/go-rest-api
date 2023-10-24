package main

import (

	"fmt"
	"github.com/mateussiilva/go-rest-api/routes"

)
const  FILE_NAME = "model.json" 



func main() {
	fmt.Println("Iniciando o servidor REST em GO")
	routes.HandleRequest()
}
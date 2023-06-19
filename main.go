package main

import (
	"github.com/RicardoArielSouza/api-go-gin/database"
	"github.com/RicardoArielSouza/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}

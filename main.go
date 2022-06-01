package main

import (
	"log"

	"github.com/otisnado/sn-api/db"
	"github.com/otisnado/sn-api/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Handlers()
}

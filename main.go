package main

import (
	"log"

	"github.com/arojas17/parqueouniajc/bd"
	"github.com/arojas17/parqueouniajc/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion")
		return
	}
	handlers.ControllerUSer()
}

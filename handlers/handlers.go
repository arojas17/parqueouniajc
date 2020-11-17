package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/arojas17/parqueouniajc/middlew"
	"github.com/arojas17/parqueouniajc/routers"
	"github.com/rs/cors"
)

/*ControllerUSer seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func ControllerUSer() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8181"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

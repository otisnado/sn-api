package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/otisnado/sn-api/middlew"
	"github.com/otisnado/sn-api/routers"
	"github.com/rs/cors"
)

/*Handlers inicializacion de servidor*/
func Handlers() {
	router := mux.NewRouter()

	/*Routes */
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

package main

import (
	"log"
	"net/http"

	v1 "github.com/guergeiro/fator-conversao-gas-portugal/cmd/v1"
	"github.com/rs/cors"
)

func main() {
	v1Router := http.NewServeMux()
	v1Router.Handle("/v1/", http.StripPrefix("/v1", v1.Routes()))

	handler := cors.Default().Handler(v1Router)

	server := http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	defer server.Close()
	log.Println("Listening on http://localhost:8000")
	log.Fatal(server.ListenAndServe())
}

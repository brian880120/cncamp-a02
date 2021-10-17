package main

import (
	"cncamp_a02/handler"
	"cncamp_a02/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handler.Healthz)
	r.HandleFunc("/notFound", handler.NotFound)
	r.HandleFunc("/badRequest", handler.BadRequest)
	r.HandleFunc("/mockSignup", handler.Signup).Methods("POST")
	r.Use(middleware.CORS)
	r.Use(middleware.ResponseHeader)
	r.Use(middleware.Log)
	http.ListenAndServe(":8090", r)
}

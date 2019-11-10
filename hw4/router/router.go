package router

import (
	"net/http"

	"github.com/Weltloose/hw4/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/"))))
	r.HandleFunc("/postForm", controller.PostFormHandler)
	r.HandleFunc("/api/unknown", controller.UnknownHandler).Methods("GET")
	r.HandleFunc("/api/testJs", controller.TestJs).Methods("GET")

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}

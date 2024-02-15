package handler

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", MakePingHandler()).Methods("GET")
	r.HandleFunc("/rating", MakeRatingHandler()).Methods("POST")
	//Erstelle eine Route für /rating
	return r
}

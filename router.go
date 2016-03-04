package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/bower_components/").Handler(http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components"))))
	router.PathPrefix("/tmpl/").Handler(http.StripPrefix("/tmpl/", http.FileServer(http.Dir("tmpl"))))
	return router
}

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handleIndexPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("index.html").Delims("*{{", "}}*").ParseFiles("pages/index.html")

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	p := page{}
	p.Prefix = "SUDDE"
	t.Delims("*{{", "}}*")
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
	}
}

package app

import (
	// "fmt"
	"net/http"
	"html/template"
)

func init() {
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Fprintf(w, "Hello world!!!!\n")
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, "hoge")
}

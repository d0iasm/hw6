package app

import (
	// "fmt"
	"html/template"
	"net/http"
)

func init() {
	http.Handle("./stylesheets/", http.StripPrefix("./stylesheets/", http.FileServer(http.Dir("./stylesheets/"))))
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Fprintf(w, "Hello world!!!!\n")
	result := "hoge"
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, result)
}

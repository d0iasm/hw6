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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Fprintf(w, "Hello world!!!!\n")
	a := r.FormValue("a")
	b := r.FormValue("b")
	result := combine(a, b)
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, result)
}

func combine(a, b string) string {
	result := ""
	shorter := ""
	longer := ""
	if len(a) < len(b) {
		shorter = a
		longer = b
	} else {
		shorter = b
		longer = a
	}

	for i := 0; i < len(shorter); i++ {
		result += longer[i : i+1]
		result += shorter[i : i+1]
	}
	for i := len(shorter); i < len(longer); i++ {
		result += longer[i:]
	}
	return result
}

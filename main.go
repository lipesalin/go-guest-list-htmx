package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Guest struct {
	Name     string
	Document string
}

func main() {
	// [TEMPLATE] - Define file template
	templ := template.Must(template.ParseFiles("index.html"))

	// [RENDER HTMX]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Add guests
		guests := map[string][]Guest{
			"Guests": {
				{Name: "Henrique", Document: "0346646546"},
				{Name: "Pedro", Document: "165651961661"},
				{Name: "Felipe", Document: "6469616161"},
			},
		}

		templ.Execute(w, guests)
	})

	// [ADD GUEST]
	http.HandleFunc("/add-guest", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		// Get params from form
		name := r.PostFormValue("name")
		document := r.PostFormValue("document")

		templ.ExecuteTemplate(w, "guest-list-element", Guest{Name: name, Document: document})
	})

	log.Fatal(http.ListenAndServe(":3131", nil))
}

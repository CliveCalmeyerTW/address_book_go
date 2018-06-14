package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Address book coming soon"))
	})

	http.ListenAndServe(":8080", r)
}
package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Address book coming soon"))
	})

	r.Route("/addresses", func(r chi.Router) {
		r.Get("/", listAddresses)
		r.Get("/search/{query}", findAddresses)
		r.Post("/", createAddress)
		r.Route("/{id:\\d+}", func(r chi.Router) {
			r.Get("/", retrieveAddress)
			r.Put("/", updateAddress)
			r.Delete("/", deleteAddress)
		})
	})

	http.ListenAndServe(":8080", r)
}

func listAddresses(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("a list of addresses"))
}

func findAddresses(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("find addresses"))
}

func createAddress(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("create an address"))
}

func retrieveAddress(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("retrieve an address"))
}

func deleteAddress(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete an address"))
}

func updateAddress(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("update an address"))
}

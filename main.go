package main

import (
	"fmt"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCxn() *sql.DB {
	dsn := "postgres://addy:addypass@localhost/address_book?sslmode=disable"
	cxn, err := sql.Open("postgres", dsn)
	checkErr(err)

	return cxn
}

func closeCxn(cxn *sql.DB) {
	cxn.Close()
}

func listAddresses(res http.ResponseWriter, req *http.Request) {
	cxn := getCxn()
	defer closeCxn(cxn)

	result := strings.Builder{}

	addresses, err := cxn.Query("SELECT id, first_name, last_name FROM address_book")
	checkErr(err)

	for addresses.Next() {
		var (
			id       int
			firsName string
			lastName string
		)
		err := addresses.Scan(&id, &firsName, &lastName)
		checkErr(err)
		fmt.Fprintf(&result, "(%d) %s %s\n", id, firsName, lastName)
	}

	res.Write([]byte(result.String()))
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

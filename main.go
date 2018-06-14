package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
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
	r.Use(render.SetContentType(render.ContentTypeJSON))

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

type AddressListItem struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *AddressListItem) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func listAddresses(res http.ResponseWriter, req *http.Request) {
	cxn := getCxn()
	defer cxn.Close()

	list := []render.Renderer{}

	addresses, err := cxn.Query(`
		SELECT 		id, 
					first_name, 
					last_name 
		FROM 		address_book
		ORDER BY 	last_name ASC
	`)
	checkErr(err)
	defer addresses.Close()

	for addresses.Next() {
		var address = new(AddressListItem)
		err := addresses.Scan(&address.Id, &address.FirstName, &address.LastName)
		checkErr(err)
		list = append(list, address)
	}

	err = render.RenderList(res, req, list)
	checkErr(err)
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

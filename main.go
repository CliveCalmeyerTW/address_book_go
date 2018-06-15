package main

import (
	"github.com/CliveCalmeyerTW/address_book_go/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"

	"./repository"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Address book coming soon"))
	})

	router.Route("/addresses", func(router chi.Router) {
		router.Get("/", listAddresses)
		// router.Get("/search/{query}", findAddresses)
		router.Post("/", createAddress)
		router.Route("/{id:\\d+}", func(router chi.Router) {
			router.Get("/", retrieveAddress)
			router.Put("/", updateAddress)
			router.Delete("/", deleteAddress)
		})
	})

	http.ListenAndServe(":8080", router)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type NotFoundError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *NotFoundError) Render(response http.ResponseWriter, request *http.Request) error {
	e.Status = 404
	e.Message = "not found"
	return nil
}

func listAddresses(response http.ResponseWriter, request *http.Request) {
	addresses, err := repository.List()
	checkErr(err)
	renderers := []render.Renderer{}

	for _, address := range addresses {
		renderers = append(renderers, &entity.AddressResponse{Address: address})
	}

	err = render.RenderList(response, request, renderers)
	checkErr(err)
}

// func findAddresses(response http.ResponseWriter, request *http.Request) {
// 	response.Write([]byte("find addresses"))
// }

func createAddress(response http.ResponseWriter, request *http.Request) {
	addressRequest := &entity.AddressRequest{}
	err := render.Bind(request, addressRequest)
	checkErr(err)
	address := addressRequest.Address
	repository.Create(address)
	render.Status(request, http.StatusCreated)
	render.Render(response, request, &entity.AddressResponse{Address: address})
}

func retrieveAddress(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	address, err := repository.Retrieve(id)
	if err != nil {
		render.Status(request, http.StatusNotFound)
		render.Render(response, request, &NotFoundError{})
		return
	}
	renderer := &entity.AddressResponse{Address: address}
	err = render.Render(response, request, renderer)
	checkErr(err)
}

func deleteAddress(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("delete an address"))
}

func updateAddress(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("update an address"))
}

package entity

import (
	"net/http"
)

type Address struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	Postcode  string `json:"postcode"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type AddressResponse struct {
	*Address
}

func (a *AddressResponse) Render(response http.ResponseWriter, request *http.Request) error {
	return nil
}

type AddressRequest struct {
	*Address
}

func (a *AddressRequest) Bind(request *http.Request) error {
	return nil
}

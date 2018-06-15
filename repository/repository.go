package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"

	"github.com/CliveCalmeyerTW/address_book_go/entity"
)

func getCxn() *sql.DB {
	dsn := "postgres://addy:addypass@localhost/address_book?sslmode=disable"
	cxn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return cxn
}

func List() ([]*entity.Address, error) {
	cxn := getCxn()
	defer cxn.Close()

	list := []*entity.Address{}

	rows, err := cxn.Query(`
        SELECT   id, 
                 first_name, 
                 last_name,
                 address_1,
                 address_2,
                 city,
                 postcode,
                 email,
                 telephone
        FROM     address_book
        ORDER BY last_name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address = new(entity.Address)
		err := rows.Scan(
			&address.Id,
			&address.FirstName,
			&address.LastName,
			&address.Address1,
			&address.Address2,
			&address.City,
			&address.Postcode,
			&address.Email,
			&address.Telephone)

		if err != nil {
			return nil, err
		}

		list = append(list, address)
	}

	return list, nil
}

func Retrieve(id string) (*entity.Address, error) {
	cxn := getCxn()
	defer cxn.Close()

	var address = new(entity.Address)
	err := cxn.QueryRow(`
		SELECT   id, 
                 first_name, 
                 last_name,
                 address_1,
                 address_2,
                 city,
                 postcode,
                 email,
                 telephone
        FROM     address_book
        WHERE    id = $1
	`, id).Scan(
		&address.Id,
		&address.FirstName,
		&address.LastName,
		&address.Address1,
		&address.Address2,
		&address.City,
		&address.Postcode,
		&address.Email,
		&address.Telephone)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func Create(address *entity.Address) error {
	cxn := getCxn()
	defer cxn.Close()

	var id string
	err := cxn.QueryRow(`
		INSERT INTO address_book
		(first_name, last_name, address_1, address_2, city, postcode, email, telephone)
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`,
		address.FirstName,
		address.LastName,
		address.Address1,
		address.Address2,
		address.City,
		address.Postcode,
		address.Email,
		address.Telephone).Scan(&id)

	if err != nil {
		return err
	}

	address.Id, err = strconv.Atoi(id)

	if err != nil {
		return err
	}
	return nil
}

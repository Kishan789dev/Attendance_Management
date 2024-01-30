package main

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID `json :"id"`
	Name     string    `json :"name"`
	Quantity int       `json :"quantity"`
}

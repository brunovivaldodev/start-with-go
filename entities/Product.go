package entities

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID string

	Name string `json:"name"`
}

type Products struct {
	Product []Product
}

func (p *Products) Add(product Product) {
	p.Product = append(p.Product, product)
}

func NewProduct(product string) *Product {
	return &Product{
		ID:   uuid.NewV4().String(),
		Name: product,
	}
}

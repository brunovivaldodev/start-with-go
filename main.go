package main

import (
	"errors"
	"fmt"

	"github.com/brunovivaldodev/strat-with-go/data"
	"github.com/brunovivaldodev/strat-with-go/entities"
	"github.com/brunovivaldodev/strat-with-go/webserver"
)

func main() {
	b := "Fusca"

	fmt.Println(b)

	var v, err = soma(-7, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(v)
	product := entities.Product{ID: "2332", Name: "Bruno Vivaldo"}
	products := entities.Products{}

	product2 := entities.NewProduct("Bruno Telafd")

	products.Add(product)

	fmt.Println(product)
	fmt.Println(products)
	fmt.Println(product2)
	data.LoadData()
	web := webserver.NewWebServer()

	web.Serve()

}

func soma(a int, b int) (int, error) {
	if a < 0 {
		return 0, errors.New("a<0")
	}
	return a + b, nil
}

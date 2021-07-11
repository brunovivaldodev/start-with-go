package data

import "github.com/brunovivaldodev/strat-with-go/entities"

var Products entities.Products

func LoadData() {
	product1 := entities.NewProduct("vivaldo")
	product2 := entities.NewProduct("tela")

	Products.Add(*product1)
	Products.Add(*product2)

}

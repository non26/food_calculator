package order

import "math/big"

type Order struct {
	Name          string
	UnitPrice     *big.Float
	NumberOfItems int
}

func NewOrder(name string, unitPrice *big.Float, numberOfItems int) Order {
	return Order{
		Name:          name,
		UnitPrice:     unitPrice,
		NumberOfItems: numberOfItems,
	}
}

package order

import "math/big"

type Orders struct {
	Orders     map[string]Order
	HaveMember bool

	totalPrice *big.Float
}

func (o *Orders) GetTotal() *big.Float {
	if o.totalPrice != nil {
		return o.totalPrice
	}

	if len(o.Orders) == 0 {
		return big.NewFloat(0)
	}

	total := big.NewFloat(0)
	for _, order := range o.Orders {
		total = new(big.Float).Add(total, new(big.Float).Mul(order.UnitPrice, big.NewFloat(float64(order.NumberOfItems))))
	}
	o.totalPrice = total
	return total
}

func (o Orders) IsMember() bool {
	return o.HaveMember
}

func (o Orders) GetOrderBy(key string) (Order, bool) {
	order, ok := o.Orders[key]
	return order, ok
}

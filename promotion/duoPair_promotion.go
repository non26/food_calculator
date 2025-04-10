package promotion

import (
	"foodcalculator/order"
	"math/big"
)

type duoPairPromotion struct {
	duoPair []string
}

func NewDuoPairPromotion() IPromotion {
	return &duoPairPromotion{
		duoPair: []string{"orange_set", "pink_set", "green_set"},
	}
}

func (p *duoPairPromotion) ValidatePromotion(orders *order.Orders) bool {
	for _, pair := range p.duoPair {
		_, ok := orders.GetOrderBy(pair)
		if !ok {
			continue
		} else {
			return true
		}
	}
	return false
}

func (p *duoPairPromotion) GetDiscount(orders *order.Orders) (discount *big.Float) {
	totalDiscount := big.NewFloat(0)
	for _, pair := range p.duoPair {
		order, ok := orders.GetOrderBy(pair)
		if !ok {
			continue
		}
		discount := p.getBundleDiscount(order)
		totalDiscount = new(big.Float).Add(totalDiscount, discount)
	}

	return totalDiscount
}

func (p *duoPairPromotion) getBundleDiscount(order order.Order) (discount *big.Float) {
	if order.NumberOfItems == 0 || order.NumberOfItems == 1 {
		return big.NewFloat(0)
	}
	var numberOfItemsInPairs *big.Float
	discountPercentage := big.NewFloat(0.05)
	quotient := order.NumberOfItems % 2
	if quotient == 1 {
		numberOfItemsInPairs = new(big.Float).SetFloat64(float64(order.NumberOfItems - 1))
	} else {
		numberOfItemsInPairs = new(big.Float).SetFloat64(float64(order.NumberOfItems))
	}
	totalPriceOfPairs := new(big.Float).Mul(order.UnitPrice, numberOfItemsInPairs)
	return new(big.Float).Mul(totalPriceOfPairs, discountPercentage)
}

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

func (p *duoPairPromotion) GetDiscount(orders *order.Orders) *big.Float {
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

func (p *duoPairPromotion) getBundleDiscount(order order.Order) *big.Float {
	pair := big.NewInt(2)
	numberOfPairs := new(big.Int).Quo(big.NewInt(int64(order.NumberOfItems)), pair)
	if numberOfPairs.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}
	discountPercentage := big.NewFloat(0.05)
	numberOfItemsInPairs := new(big.Float).SetFloat64(float64(numberOfPairs.Int64() * pair.Int64()))
	totalPriceOfPairs := new(big.Float).Mul(order.UnitPrice, numberOfItemsInPairs)
	return new(big.Float).Mul(totalPriceOfPairs, discountPercentage)
}

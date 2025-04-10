package calculator

import (
	"foodcalculator/order"
	"foodcalculator/promotion"
	"math/big"
)

type ICalculator interface {
	Calculate(Order *order.Orders) (finalPrice *big.Float)
}

type Calculator struct {
	memberPromotion  promotion.IPromotion
	duoPairPromotion promotion.IPromotion
}

func NewCalculator(memberPromotion promotion.IPromotion, duoPairPromotion promotion.IPromotion) ICalculator {
	return &Calculator{
		memberPromotion:  memberPromotion,
		duoPairPromotion: duoPairPromotion,
	}
}

func (d *Calculator) Calculate(Order *order.Orders) (finalPrice *big.Float) {

	totalDiscount := big.NewFloat(0)
	if d.memberPromotion.ValidatePromotion(Order) {
		discount := d.memberPromotion.GetDiscount(Order)
		totalDiscount = new(big.Float).Add(totalDiscount, discount)
	}
	if d.duoPairPromotion.ValidatePromotion(Order) {
		discount := d.duoPairPromotion.GetDiscount(Order)
		totalDiscount = new(big.Float).Add(totalDiscount, discount)
	}

	total := Order.GetTotal()
	if totalDiscount.Cmp(big.NewFloat(0)) == 0 {
		return total
	}
	finalPrice = new(big.Float).Sub(total, totalDiscount)
	return finalPrice
}

package promotion

import (
	"foodcalculator/order"
	"math/big"
)

type memberPromotion struct {
}

func NewMemberPromotion() IPromotion {
	return &memberPromotion{}
}

func (p *memberPromotion) ValidatePromotion(orders *order.Orders) bool {
	return orders.IsMember()
}

func (p *memberPromotion) GetDiscount(orders *order.Orders) *big.Float {
	total := orders.GetTotal()
	discount := new(big.Float).Mul(total, big.NewFloat(0.1))
	return discount
}

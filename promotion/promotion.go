package promotion

import (
	"foodcalculator/order"
	"math/big"
)

type IPromotion interface {
	ValidatePromotion(orders *order.Orders) bool
	GetDiscount(orders *order.Orders) *big.Float
}

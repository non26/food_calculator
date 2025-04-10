package promotion_test

import (
	"foodcalculator/order"
	"foodcalculator/promotion"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var RedSet = "red_set"
var GreenSet = "green_set"
var BlueSet = "blue_set"
var YellowSet = "yellow_set"
var PinkSet = "pink_set"
var PurpleSet = "purple_set"
var OrangeSet = "orange_set"

func TestMemberPromotion_ValidatePromotion(t *testing.T) {
	promotion := promotion.NewMemberPromotion()
	t.Run("test validate promotion with member", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: true,
		}

		// Action
		result := promotion.ValidatePromotion(orders)

		// Assert
		assert.True(t, result)
	})

	t.Run("test validate promotion with no member", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}

		// Action
		result := promotion.ValidatePromotion(orders)

		// Assert
		assert.False(t, result)
	})

	t.Run("test get discount with member", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: true,
		}
		orders.Orders[RedSet] = order.NewOrder(RedSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedDiscount := big.NewFloat(10)
		// Action
		totalDiscount := promotion.GetDiscount(orders)
		// Assert
		assert.Equal(t, expectedDiscount.String(), totalDiscount.String())
	})

	t.Run("test get discount with member and decimal point of price", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: true,
		}
		orders.Orders[RedSet] = order.NewOrder(RedSet, big.NewFloat(10.25), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedDiscount := big.NewFloat(10.025)
		// Action
		totalDiscount := promotion.GetDiscount(orders)
		// Assert
		assert.Equal(t, expectedDiscount.String(), totalDiscount.String())
	})
}

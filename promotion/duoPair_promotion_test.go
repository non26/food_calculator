package promotion_test

import (
	"foodcalculator/order"
	"foodcalculator/promotion"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuoPairPromotion_ValidatePromotion(t *testing.T) {
	promotion := promotion.NewDuoPairPromotion()
	t.Run("test validate promotion with duo pair", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[GreenSet] = order.NewOrder(GreenSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		// Action
		actual := promotion.ValidatePromotion(orders)
		// Assert
		assert.True(t, actual)
	})

	t.Run("test validate promotion with not duo pair", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[RedSet] = order.NewOrder(RedSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		// Action
		actual := promotion.ValidatePromotion(orders)
		// Assert
		assert.False(t, actual)
	})

	t.Run("test get discount with not enough duo pair", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[GreenSet] = order.NewOrder(GreenSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedDiscount := big.NewFloat(0)
		// Action
		actual := promotion.GetDiscount(orders)
		// Assert
		assert.Equal(t, expectedDiscount.String(), actual.String())
	})

	t.Run("test get discount with enough duo pairs of 1 out of 2 items", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[GreenSet] = order.NewOrder(GreenSet, big.NewFloat(10), 2)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedDiscount := big.NewFloat(1)
		// Action
		actual := promotion.GetDiscount(orders)
		// Assert
		assert.Equal(t, expectedDiscount.String(), actual.String())
	})

	t.Run("test get discount with enough duo pairs of 1 out of 3 items", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[GreenSet] = order.NewOrder(GreenSet, big.NewFloat(10), 3)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedDiscount := big.NewFloat(1)
		// Action
		actual := promotion.GetDiscount(orders)
		// Assert
		assert.Equal(t, expectedDiscount.String(), actual.String())
	})

	t.Run("test get discount with enough duo pairs of 2 out of 4 items", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[GreenSet] = order.NewOrder(GreenSet, big.NewFloat(10), 4)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedDiscount := big.NewFloat(2)
		// Action
		actual := promotion.GetDiscount(orders)
		// Assert
		assert.Equal(t, expectedDiscount.String(), actual.String())
	})

}

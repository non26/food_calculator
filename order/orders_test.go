package order_test

import (
	"foodcalculator/order"
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

func TestOrders_GetTotal(t *testing.T) {
	t.Run("test get total with no orders", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}

		// Action
		total := orders.GetTotal()

		// Assert
		assert.Equal(t, big.NewFloat(0).String(), total.String())
	})

	t.Run("test get total with orders	", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[RedSet] = order.NewOrder(RedSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)

		// Action
		total := orders.GetTotal()

		// Assert
		assert.Equal(t, big.NewFloat(100).String(), total.String())
	})
}

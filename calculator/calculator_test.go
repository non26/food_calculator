package calculator_test

import (
	"foodcalculator/calculator"
	"foodcalculator/mocks"
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

func TestCalculator_Calculate(t *testing.T) {

	t.Run("test calculate without discount", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[RedSet] = order.NewOrder(RedSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedFinalPrice := big.NewFloat(100)
		memberPromotion := mocks.NewIPromotion(t)
		memberPromotion.On("ValidatePromotion", orders).Return(false)
		duoPairPromotion := mocks.NewIPromotion(t)
		duoPairPromotion.On("ValidatePromotion", orders).Return(false)
		calculator := calculator.NewCalculator(memberPromotion, duoPairPromotion)

		// Action
		actual := calculator.Calculate(orders)

		// Assert
		assert.Equal(t, expectedFinalPrice.String(), actual.String())
	})

	t.Run("test calculate with only member discount", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: true,
		}
		orders.Orders[RedSet] = order.NewOrder(RedSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)
		expectedFinalPrice := big.NewFloat(90)

		memberPromotion := mocks.NewIPromotion(t)
		memberPromotion.On("ValidatePromotion", orders).Return(true)
		memberPromotion.On("GetDiscount", orders).Return(big.NewFloat(10))

		duoPairPromotion := mocks.NewIPromotion(t)
		duoPairPromotion.On("ValidatePromotion", orders).Return(false)

		calculator := calculator.NewCalculator(memberPromotion, duoPairPromotion)

		// Action
		actual := calculator.Calculate(orders)

		// Assert
		assert.Equal(t, expectedFinalPrice.String(), actual.String())
	})

	t.Run("test calculate with only duo pair discount", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[OrangeSet] = order.NewOrder(OrangeSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)

		expectedFinalPrice := big.NewFloat(99)

		memberPromotion := mocks.NewIPromotion(t)
		memberPromotion.On("ValidatePromotion", orders).Return(false)

		duoPairPromotion := mocks.NewIPromotion(t)
		duoPairPromotion.On("ValidatePromotion", orders).Return(true)
		duoPairPromotion.On("GetDiscount", orders).Return(big.NewFloat(1))

		calculator := calculator.NewCalculator(memberPromotion, duoPairPromotion)

		// Action
		actual := calculator.Calculate(orders)

		// Assert
		assert.Equal(t, expectedFinalPrice.String(), actual.String())
	})

	t.Run("test calculate with both discounts", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[OrangeSet] = order.NewOrder(OrangeSet, big.NewFloat(10), 1)
		orders.Orders[BlueSet] = order.NewOrder(BlueSet, big.NewFloat(20), 2)
		orders.Orders[YellowSet] = order.NewOrder(YellowSet, big.NewFloat(25), 2)

		expectedFinalPrice := big.NewFloat(89)

		memberPromotion := mocks.NewIPromotion(t)
		memberPromotion.On("ValidatePromotion", orders).Return(true)
		memberPromotion.On("GetDiscount", orders).Return(big.NewFloat(10))

		duoPairPromotion := mocks.NewIPromotion(t)
		duoPairPromotion.On("ValidatePromotion", orders).Return(true)
		duoPairPromotion.On("GetDiscount", orders).Return(big.NewFloat(1))

		calculator := calculator.NewCalculator(memberPromotion, duoPairPromotion)

		// Action
		actual := calculator.Calculate(orders)

		// Assert
		assert.Equal(t, expectedFinalPrice.String(), actual.String())
	})

}

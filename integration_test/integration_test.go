package integrationtest_test

import (
	"foodcalculator/calculator"
	"foodcalculator/order"
	"foodcalculator/promotion"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var redSet = "red_set"
var greenSet = "green_set"
var blueSet = "blue_set"
var yellowSet = "yellow_set"
var pinkSet = "pink_set"
var purpleSet = "purple_set"
var orangeSet = "orange_set"

type Set struct {
	Name  string
	Price *big.Float
}

var RedSet = Set{
	Name:  redSet,
	Price: big.NewFloat(50),
}

var GreenSet = Set{
	Name:  greenSet,
	Price: big.NewFloat(40),
}

var BlueSet = Set{
	Name:  blueSet,
	Price: big.NewFloat(30),
}

var YellowSet = Set{
	Name:  yellowSet,
	Price: big.NewFloat(50),
}

var PinkSet = Set{
	Name:  pinkSet,
	Price: big.NewFloat(80),
}

var PurpleSet = Set{
	Name:  purpleSet,
	Price: big.NewFloat(90),
}

var OrangeSet = Set{
	Name:  orangeSet,
	Price: big.NewFloat(120),
}

func TestCalculate(t *testing.T) {
	memberPromotion := promotion.NewMemberPromotion()
	duoPairPromotion := promotion.NewDuoPairPromotion()
	calculator := calculator.NewCalculator(memberPromotion, duoPairPromotion)

	t.Run("test calculate with no dicount", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[RedSet.Name] = order.NewOrder(RedSet.Name, RedSet.Price, 1)
		orders.Orders[BlueSet.Name] = order.NewOrder(BlueSet.Name, BlueSet.Price, 2)
		orders.Orders[YellowSet.Name] = order.NewOrder(YellowSet.Name, YellowSet.Price, 2)
		expectedFinalPrice := big.NewFloat(210)
		// Action
		total := calculator.Calculate(orders)
		// Assert
		assert.Equal(t, expectedFinalPrice.String(), total.String())
	})

	t.Run("test calculate with only member discount", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: true,
		}
		orders.Orders[RedSet.Name] = order.NewOrder(RedSet.Name, RedSet.Price, 1)
		orders.Orders[BlueSet.Name] = order.NewOrder(BlueSet.Name, BlueSet.Price, 2)
		orders.Orders[YellowSet.Name] = order.NewOrder(YellowSet.Name, YellowSet.Price, 2)
		expectedFinalPrice := big.NewFloat(189)
		// Action
		total := calculator.Calculate(orders)
		// Assert
		assert.Equal(t, expectedFinalPrice.String(), total.String())
	})

	t.Run("test calculate with only duo pair discount of 2 pair per item", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[RedSet.Name] = order.NewOrder(RedSet.Name, RedSet.Price, 1)
		orders.Orders[OrangeSet.Name] = order.NewOrder(OrangeSet.Name, OrangeSet.Price, 2)
		orders.Orders[PinkSet.Name] = order.NewOrder(PinkSet.Name, PinkSet.Price, 2)
		expectedFinalPrice := big.NewFloat(430)
		// Action
		total := calculator.Calculate(orders)
		// Assert
		assert.Equal(t, expectedFinalPrice.String(), total.String())
	})

	t.Run("test calculate with only duo pair discount of 2 pairs per item with exceed item", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: false,
		}
		orders.Orders[RedSet.Name] = order.NewOrder(RedSet.Name, RedSet.Price, 1)
		orders.Orders[OrangeSet.Name] = order.NewOrder(OrangeSet.Name, OrangeSet.Price, 3)
		orders.Orders[PinkSet.Name] = order.NewOrder(PinkSet.Name, PinkSet.Price, 2)
		expectedFinalPrice := big.NewFloat(550)
		// Action
		total := calculator.Calculate(orders)
		// Assert
		assert.Equal(t, expectedFinalPrice.String(), total.String())
	})

	t.Run("test calculate with both discount", func(t *testing.T) {
		// Arrangement
		orders := &order.Orders{
			Orders:     map[string]order.Order{},
			HaveMember: true,
		}
		orders.Orders[RedSet.Name] = order.NewOrder(RedSet.Name, RedSet.Price, 1)
		orders.Orders[OrangeSet.Name] = order.NewOrder(OrangeSet.Name, OrangeSet.Price, 2)
		orders.Orders[PinkSet.Name] = order.NewOrder(PinkSet.Name, PinkSet.Price, 2)
		expectedFinalPrice := big.NewFloat(385)
		// Action
		total := calculator.Calculate(orders)
		// Assert
		assert.Equal(t, expectedFinalPrice.String(), total.String())
	})
}

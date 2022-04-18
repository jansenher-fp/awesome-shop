package integration_test

import (
	"awesome-shop/cart"
	"awesome-shop/inventory"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stocks = map[string]int{
		"loofah":  5,
		"soap":    10,
		"shampoo": 3,
	}

	priceList = map[string]int{
		"loofah":  5,
		"soap":    1,
		"shampoo": 2,
	}
)

func TestCartInventoryIntegration(t *testing.T) {
	inventoryService := inventory.NewService(stocks, priceList)
	cartService := cart.NewService(inventoryService)
	userToken := "testing"

	err := cartService.AddItem(userToken, "loofah", 3)
	cart := cartService.GetCart(userToken)

	assert.NoError(t, err)
	assert.Equal(t, map[string]int{"loofah": 3}, cart.Items())

	err = cartService.AddItem(userToken, "loofah", 1)
	cart = cartService.GetCart(userToken)
	assert.Equal(t, map[string]int{"loofah": 4}, cart.Items())

	stocksRemaining := inventoryService.GetInventory()
	assert.Equal(t, map[string]int{
		"loofah":  1,
		"soap":    10,
		"shampoo": 3,
	}, stocksRemaining)
}

func TestCartInventoryErrors(t *testing.T) {
	inventoryService := inventory.NewService(stocks, priceList)
	cartService := cart.NewService(inventoryService)
	userToken := "testing"

	err := cartService.AddItem(userToken, "shampoo", 10)
	cart := cartService.GetCart(userToken)

	assert.ErrorContains(t, err, fmt.Sprintf("insufficient stocks for %s, only %d qty left", "shampoo", 3))
	assert.Equal(t, map[string]int{}, cart.Items())

	stocksRemaining := inventoryService.GetInventory()
	assert.Equal(t, stocks, stocksRemaining)
}

package integration_test

import (
	"awesome-shop/cart"
	"awesome-shop/inventory"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartInventoryIntegration(t *testing.T) {
	stocks := map[string]int{
		"loofah":  5,
		"soap":    10,
		"shampoo": 3,
	}

	inventoryService := inventory.NewService(stocks)
	cartService := cart.NewService(inventoryService)

	err := cartService.AddItem("loofah", 3)
	cart := cartService.GetCart()

	assert.NoError(t, err)
	assert.Equal(t, map[string]int{"loofah": 3}, cart)

	err = cartService.AddItem("loofah", 1)
	cart = cartService.GetCart()
	assert.Equal(t, map[string]int{"loofah": 4}, cart)

	stocksRemaining := inventoryService.GetInventory()
	assert.Equal(t, map[string]int{
		"loofah":  1,
		"soap":    10,
		"shampoo": 3,
	}, stocksRemaining)
}

func TestCartInventoryErrors(t *testing.T) {
	stocks := map[string]int{
		"loofah":  5,
		"soap":    10,
		"shampoo": 3,
	}

	inventoryService := inventory.NewService(stocks)
	cartService := cart.NewService(inventoryService)

	err := cartService.AddItem("shampoo", 10)
	cart := cartService.GetCart()

	assert.ErrorContains(t, err, fmt.Sprintf("insufficient stocks for %s, only %d qty left", "shampoo", 3))
	assert.Equal(t, map[string]int{}, cart)

	stocksRemaining := inventoryService.GetInventory()
	assert.Equal(t, stocks, stocksRemaining)
}

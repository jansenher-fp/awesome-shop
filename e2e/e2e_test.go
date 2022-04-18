package e2e_test

import (
	"awesome-shop/cart"
	"awesome-shop/checkout"
	"awesome-shop/inventory"
	"awesome-shop/shop"
	"awesome-shop/user"
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

func TestHappyFlow(t *testing.T) {
	checkoutService := checkout.NewService()
	inventory := inventory.NewService(stocks, priceList)
	userService := user.NewUserService(user.AnyLogin(user.AllLogin), user.NewInMemoryUserDB())
	cartService := cart.NewService(inventory)

	shop := shop.NewShop(checkoutService, cartService, userService, inventory)
	// User log in flow
	t.Logf("Step %d: Adding user", 0)
	err := shop.AddUser("John", "mypassword", "987876")
	assert.NoError(t, err, "no error expected in adding new user")

	t.Logf("Step %d: User log in successful", 1)
	token, err := shop.Login("John", "mypassword")
	assert.NoError(t, err, "no error expected in logging in user")

	// Cart flow
	t.Logf("Step %d: Adding item to cart", 2)
	err = shop.AddItem(token, "soap", 2)
	assert.NoError(t, err, "no error expected in adding soap to cart")

	t.Logf("Step %d: Adding more items to cart", 3)
	err = shop.AddItem(token, "shampoo", 1)
	assert.NoError(t, err, "no error expected in adding shampoo to cart")

	// Checkout flow
	t.Logf("Step %d: Checking out cart", 4)
	cardDetails := "demo card details"
	status, err := shop.Checkout(token, cardDetails)
	assert.Equal(t, checkout.Successful, status)

	t.Logf("Step %d: Cart checkout success: user history updated", 5)
	testUser, err := shop.GetUser(token)
	assert.NoError(t, err, "no error expected getting user")

	userHx := testUser.GetHistory()
	expectedItems := map[string]int{
		"soap":    2,
		"shampoo": 1,
	}
	allHistoryItems := make([]map[string]int, len(userHx))

	for _, order := range userHx {
		allHistoryItems = append(allHistoryItems, order.GetItems())
	}
	assert.Contains(t, allHistoryItems, expectedItems)

	t.Logf("Step %d: Cart checkout success: inventory updated", 6)
	expectedInventory := map[string]int{
		"loofah":  5,
		"soap":    8,
		"shampoo": 2,
	}
	assert.Equal(t, expectedInventory, shop.GetInventory())
}

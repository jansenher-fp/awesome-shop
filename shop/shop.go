package shop

import (
	"awesome-shop/cart"
	"awesome-shop/checkout"
	"awesome-shop/inventory"
	"awesome-shop/user"
	"fmt"
)

type Shop struct {
	checkoutService *checkout.Service
	userService     *user.Service
	cartService     *cart.Service
	inventory       *inventory.Service
}

func NewShop(
	checkoutService *checkout.Service,
	cartService *cart.Service,
	userService *user.Service,
	inventory *inventory.Service,
) *Shop {
	return &Shop{
		checkoutService: checkoutService,
		userService:     userService,
		cartService:     cartService,
		inventory:       inventory,
	}
}

func (s *Shop) AddUser(name, password, phone string) error {
	return s.userService.AddUser(name, password, phone)
}

func (s *Shop) Login(name, password string) (string, error) {
	return s.userService.Login(name, password)
}

func (s *Shop) GetUser(token string) (*user.User, error) {
	return s.userService.GetUser(token)
}

func (s *Shop) AddItem(token, name string, qty int) error {
	return s.cartService.AddItem(token, name, qty)
}

func (s *Shop) GetInventory() map[string]int {
	return s.inventory.GetInventory()
}

func (s *Shop) Checkout(token, cardDetails string) (string, error) {
	userCart := s.cartService.GetCart(token)
	u, err := s.userService.GetUser(token)

	if err != nil {
		return "", fmt.Errorf("checkout failed: %w", err)
	}

	status, err := s.checkoutService.Checkout(userCart.Amount(), cardDetails)

	if err != nil {
		return "", fmt.Errorf("checkout failed: %w", err)
	}
	u.AddHistory(user.NewOrder(userCart.Items()))

	return status, nil
}

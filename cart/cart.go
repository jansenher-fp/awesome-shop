package cart

import (
	"fmt"
)

type inventoryService interface {
	GetStock(string) int
	GetPrice(string) int
	DeductStock(string, int) error
}

type Cart struct {
	userToken string
	items     map[string]int
	amount    int
}

type Service struct {
	inventory inventoryService
	carts     map[string]*Cart
}

func NewService(inventory inventoryService) *Service {
	return &Service{
		inventory: inventory,
		carts:     make(map[string]*Cart),
	}
}

func (s *Service) AddItem(token, name string, qty int) error {
	stock := s.inventory.GetStock(name)
	if stock < qty {
		return fmt.Errorf("insufficient stocks for %s, only %d qty left", name, stock)
	}

	if err := s.inventory.DeductStock(name, qty); err != nil {
		return fmt.Errorf("inventory error: %w", err)
	}

	userCart := s.GetCart(token)

	userCart.items[name] += qty

	price := s.inventory.GetPrice(name)
	userCart.amount += price * qty

	return nil
}

func (s *Service) GetCart(token string) *Cart {
	return s.getOrCreateCart(token)
}

func (c *Cart) Items() map[string]int {
	return c.items
}

func (c *Cart) Amount() int {
	return c.amount
}

func (s *Service) getOrCreateCart(token string) *Cart {
	userCart, ok := s.carts[token]
	if !ok {
		s.carts[token] = &Cart{
			userToken: token,
			items:     make(map[string]int),
		}
		userCart = s.carts[token]
	}
	return userCart
}

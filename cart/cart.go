package cart

import (
	"fmt"
)

type inventoryService interface {
	GetStock(string) int
	DeductStock(string, int) error
}

type cart struct {
	id    string
	items map[string]int
}

type Service struct {
	inventory inventoryService
	carts     map[string]cart
}

func NewService(inventory inventoryService) *Service {
	return &Service{
		inventory: inventory,
		carts:     make(map[string]cart),
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

	userCart, ok := s.carts[token]
	if !ok {
		s.carts[token] = cart{
			id:    token,
			items: make(map[string]int),
		}
		userCart = s.carts[token]
	}

	userCart.items[name] += qty

	return nil
}

func (s *Service) GetCart(token string) cart {
	return s.carts[token]
}

func (c *cart) Items() map[string]int {
	return c.items
}

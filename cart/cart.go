package cart

import (
	"fmt"
)

type inventoryService interface {
	GetStock(string) int
	DeductStock(string, int) error
}

type Service struct {
	inventory inventoryService
	cart      map[string]int
}

func NewService(inventory inventoryService) *Service {
	return &Service{
		inventory: inventory,
		cart:      make(map[string]int),
	}
}

func (s *Service) AddItem(name string, qty int) error {
	stock := s.inventory.GetStock(name)
	if stock < qty {
		return fmt.Errorf("insufficient stocks for %s, only %d qty left", name, stock)
	}

	if err := s.inventory.DeductStock(name, qty); err != nil {
		return fmt.Errorf("inventory error: %w", err)
	}

	s.cart[name] += qty
	return nil
}

func (s *Service) GetCart() map[string]int {
	return s.cart
}

package inventory

import "fmt"

type Service struct {
	inventory map[string]int
	priceList map[string]int
}

func NewService(inventory, priceList map[string]int) *Service {
	return &Service{inventory: inventory, priceList: priceList}
}

func (s *Service) GetInventory() map[string]int {
	return s.inventory
}

func (s *Service) GetStock(name string) int {
	return s.inventory[name]
}

func (s *Service) DeductStock(name string, qty int) error {
	if qty > s.inventory[name] {
		return fmt.Errorf("insufficient stocks for %s", name)
	}
	s.inventory[name] -= qty
	return nil
}

func (s *Service) GetPrice(name string) int {
	return s.priceList[name]
}

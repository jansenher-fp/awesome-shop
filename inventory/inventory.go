package inventory

import "fmt"

type Service struct {
	inventory map[string]int
}

func NewService(inventory map[string]int) *Service {
	return &Service{inventory: inventory}
}

func (i *Service) GetInventory() map[string]int {
	return i.inventory
}

func (i *Service) GetStock(name string) int {
	return i.inventory[name]
}

func (i *Service) DeductStock(name string, qty int) error {
	if qty > i.inventory[name] {
		return fmt.Errorf("insufficient stocks for %s", name)
	}
	i.inventory[name] -= qty
	return nil
}

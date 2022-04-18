package checkout

import (
	"fmt"
)

const (
	Successful = "Successful"
	Incomplete = "Incomplete"
)

type Service struct {
}

func (s *Service) Verify(amount int, cardDetails string) (string, error) {
	fmt.Printf("Successfully verified card %s for payment of $%d\n", cardDetails, amount)
	return Successful, nil
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Checkout(amount int, cardDetails string) (string, error) {
	status, err := s.Verify(amount, cardDetails)
	if err != nil || status != Successful {
		return Incomplete, fmt.Errorf("unable to verify status %s with error: %w", status, err)
	}
	return status, nil
}

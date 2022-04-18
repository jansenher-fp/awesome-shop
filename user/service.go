package user

import (
	"errors"
	"fmt"
)

var UserNotFoundErr = errors.New("user not found")

type Service struct {
	authProvider authProvider
	userDatabase userDatabase
}

type authProvider interface {
	authenticate(name, password string) (string, error)
}

type userDatabase interface {
	getUser(token string) (*User, error)
	addUser(token string, user *User) error
}

func NewUserService(ap authProvider, userDatabase userDatabase) *Service {
	return &Service{authProvider: ap, userDatabase: userDatabase}
}

func (s *Service) Login(name, password string) (string, error) {
	return s.authProvider.authenticate(name, password)
}

func (s *Service) GetUser(token string) (*User, error) {
	return s.userDatabase.getUser(token)
}

func (s *Service) AddUser(name, password, phone string) error {
	token, err := s.Login(name, password)

	if err != nil {
		return fmt.Errorf("unable to add user: %w", err)
	}
	user := NewUser(name, phone)
	err = s.userDatabase.addUser(token, user)

	if err != nil {
		return fmt.Errorf("unable to add user: %w", err)
	}
	return nil
}

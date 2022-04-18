package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	AllLoginProvider = AnyLogin(AllLogin)
)

func TestLogin(t *testing.T) {
	inMemoryUserDB := NewInMemoryUserDB()
	userService := NewUserService(AllLoginProvider, inMemoryUserDB)
	token, err := userService.Login("hi", "demopassword")
	assert.Equal(t, "a53b655f", token)
	assert.NoError(t, err)
}

func TestAddandGetUser(t *testing.T) {
	inMemoryUserDB := NewInMemoryUserDB()
	userService := NewUserService(AllLoginProvider, inMemoryUserDB)

	err := userService.AddUser("Bob", "Odenkirk", "11223344")
	assert.NoError(t, err, "no error expected when adding new user")

	token, err := userService.Login("Bob", "Odenkirk")
	assert.NoError(t, err)

	user, err := userService.GetUser(token)
	expected := &User{
		name:       "Bob",
		phone:      "11223344",
		favourites: map[string]bool{},
		history:    map[string]Order{},
	}
	assert.Equal(t, expected, user, "expected %+v but got %+v", expected, user)

	err = userService.AddUser("Bob", "Odenkirk", "11223344")
	assert.ErrorIs(t, err, UserAlreadyExistsErr)
}

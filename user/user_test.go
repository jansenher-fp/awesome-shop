package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	testUser := NewUser("Alice", "88888888")
	expectedUser := &User{
		name:       "Alice",
		phone:      "88888888",
		favourites: make(map[string]bool),
		history:    make(map[string]Order),
	}
	assert.Equal(t, expectedUser, testUser)
}

func TestUserAddFavourite(t *testing.T) {
	testUser := NewUser("Alice", "88888888")
	testUser.AddFavourite("ice cream")
	testUser.AddFavourite("chocolate")

	assert.True(t, testUser.Favours("ice cream"), "expected ice cream to be in user's favourite but it is not")
	assert.True(t, testUser.Favours("chocolate"), "expected chocolate to be in user's favourite but it is not")
	assert.False(t, testUser.Favours("donuts"), "expected donuts not to be in user's favourite but it is")
}

func TestUserAddFavouriteVariation(t *testing.T) {
	testUser := NewUser("Alice", "88888888")
	testUser.AddFavourite("ice cream")
	testUser.AddFavourite("chocolate")

	/* NOTE: This would pass too, but it will be coupled to the implementation -> will have to change if there's
	refactoring
	*/
	assert.Equal(t, map[string]bool{
		"ice cream": true,
		"chocolate": true,
	}, testUser.favourites)
}

func TestUserRemoveFavourite(t *testing.T) {
	testUser := NewUser("Alice", "88888888")
	testUser.AddFavourite("ice cream")
	testUser.AddFavourite("chocolate")
	assert.True(t, testUser.Favours("ice cream"), "expected ice cream to be in user's favourite but it is not")

	testUser.RemoveFavourite("ice cream")
	assert.True(t, testUser.Favours("chocolate"), "expected chocolate to be in user's favourite but it is not")
	assert.False(t, testUser.Favours("ice cream"), "expected ice cream to be removed from user's favourite but it is still favoured")
}

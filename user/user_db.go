package user

import "errors"

type InMemoryUserDB map[string]*User

var UserAlreadyExistsErr = errors.New("user already exists")

func (db InMemoryUserDB) getUser(token string) (*User, error) {
	u, ok := db[token]
	if !ok {
		return &User{}, UserNotFoundErr
	}
	return u, nil
}

func (db InMemoryUserDB) addUser(token string, user *User) error {
	_, ok := db[token]
	if ok {
		return UserAlreadyExistsErr
	}

	db[token] = user
	return nil
}

func NewInMemoryUserDB() InMemoryUserDB {
	return make(InMemoryUserDB)
}

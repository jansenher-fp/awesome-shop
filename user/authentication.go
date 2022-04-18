package user

import (
	"crypto/md5"
	"errors"
	"fmt"
)

var UnauthorisedUserError error = errors.New("unauthorised user")

type AnyLogin func(name, password string) (string, error)

func (f AnyLogin) authenticate(name, password string) (string, error) {
	return f(name, password)
}

func AllLogin(name, password string) (string, error) {
	if name == "errorUser" {
		return "", UnauthorisedUserError
	}

	hmd5 := md5.Sum([]byte(name + password))
	return fmt.Sprintf("%x", hmd5)[:8], nil
}

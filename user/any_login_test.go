package user

import (
	"crypto/md5"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var UnauthorisedUserError error = errors.New("unauthorised user")

func AllLogin(name, password string) (string, error) {
	if name == "errorUser" {
		return "", UnauthorisedUserError
	}

	hmd5 := md5.Sum([]byte(name + password))
	return fmt.Sprintf("%x", hmd5)[:8], nil
}
func TestAnyLogin(t *testing.T) {
	userDB := AnyLogin(AllLogin)
	loginService := NewLoginService(userDB)
	token, err := loginService.Login("hi", "demopassword")
	assert.Equal(t, "a53b655f", token)
	assert.NoError(t, err)
}

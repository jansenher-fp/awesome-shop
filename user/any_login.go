package user

type AnyLogin func(name, password string) (string, error)

func (f AnyLogin) authenticate(name, password string) (string, error) {
	return f(name, password)
}

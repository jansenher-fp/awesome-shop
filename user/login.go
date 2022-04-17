package user

type LoginService struct {
	database UserDatabase
}

type UserDatabase interface {
	authenticate(name, password string) (string, error)
}

func NewLoginService(d UserDatabase) *LoginService {
	return &LoginService{database: d}
}

func (s *LoginService) Login(name, password string) (string, error) {
	return s.database.authenticate(name, password)
}

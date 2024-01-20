package user

type Service interface {
	Authenticate(username, password string) (bool, error)
}

func Authenticate(username, password string) (bool, error) {
	return false, nil
}

package user

type LoginInSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginInSchema) Validate() error {
	return nil
}

type LoginOutSchema struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

package user

import "github.com/sirrobot01/oauth-sso/internal/common"

type LoginInSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginInSchema) Validate() error {
	if login.Username == "" {
		return &common.ValidationError{
			Message: "username is required",
		}
	}
	if login.Password == "" {
		return &common.ValidationError{
			Message: "password is required",
		}
	}
	return nil
}

type RegisterInSchema struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (register *RegisterInSchema) Validate() error {
	if register.Username == "" {
		return &common.ValidationError{
			Message: "username is required",
		}
	}
	if register.Password == "" {
		return &common.ValidationError{
			Message: "password is required",
		}
	}
	if register.ConfirmPassword == "" {
		return &common.ValidationError{
			Message: "confirm password is required",
		}
	}
	if register.Password != register.ConfirmPassword {
		return &common.ValidationError{
			Message: "password and confirm password must match",
		}
	}
	return nil
}

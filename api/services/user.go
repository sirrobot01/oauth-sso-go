package services

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/api/schemas"
	"github.com/sirrobot01/oauth-sso/api/storage/db"
)

func (s *Service) RegisterUser(user *schemas.RegisterInSchema) (*models.User, error) {
	password, err := common.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	u := &models.User{
		Username: user.Username,
		Password: password,
	}
	err = db.CreateUser(s.cfg.DB, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) LoginUser(u *schemas.LoginInSchema) (*models.User, error) {
	user, err := db.GetUserByUsername(s.cfg.DB, u.Username)
	if err != nil {
		return nil, err
	}
	if !common.CheckPassword(u.Password, user.Password) {
		return nil, &common.ValidationError{
			Message: "invalid password",
		}
	}
	return &user, nil
}

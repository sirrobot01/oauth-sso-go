package services

import (
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/api/schemas"
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
	err = s.cfg.DB.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) LoginUser(u *schemas.LoginInSchema) (*models.User, error) {
	user, err := s.cfg.DB.GetUserByQuery("username = ?", u.Username)
	if err != nil {
		return nil, err
	}
	//user, ok := obj.(*models.User)
	if !common.CheckPassword(u.Password, user.Password) {
		return nil, &common.ValidationError{
			Message: "invalid password",
		}
	}
	return user, nil
}

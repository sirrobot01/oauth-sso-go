package services

import (
	"github.com/google/uuid"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/api/schemas"
)

func (s *Service) CreateApp(app *schemas.AppRequest) (*models.App, error) {
	clientKey := uuid.New().String()
	clientSecret := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(clientKey)).String()
	a := &models.App{
		Name:         app.Name,
		RedirectURIs: app.RedirectUris,
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
		Scopes:       app.Scopes,
		UserID:       app.UserID,
	}
	err := s.cfg.DB.CreateApp(a)
	if err != nil {
		return nil, err
	}
	return a, nil

}

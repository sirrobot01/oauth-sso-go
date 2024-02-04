package services

import (
	"github.com/google/uuid"
	"github.com/sirrobot01/oauth-sso/api/common"
	"github.com/sirrobot01/oauth-sso/api/models"
	"github.com/sirrobot01/oauth-sso/api/schemas"
	"strconv"
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

func (s *Service) GetClient(clientID string) (*models.App, error) {
	return s.cfg.DB.GetAppByQuery("client_key = ?", clientID)
}

func (s *Service) AuthorizeClient(app *models.App, params *schemas.AuthorizeInSchema) (*models.App, error) {
	if params.ResponseType == "" {
		params.ResponseType = "code"
	}
	if params.ResponseType == "code" {
		// generate auth code
		return s.GenerateAuthCode(app)
	} else {
		// generate token
		return s.GenerateToken(app)
	}

}

func (s *Service) GenerateAuthCode(app *models.App) (*models.App, error) {
	authCode := uuid.New().String()
	app.AuthCode = authCode
	err := s.cfg.DB.UpdateApp(app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (s *Service) GenerateToken(app *models.App) (*models.App, error) {
	token, _ := common.GenerateAppToken(strconv.Itoa(int(app.ID)), app.Name)
	// TODO: Implement refresh token etc
	app.AuthCode = token
	err := s.cfg.DB.UpdateApp(app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

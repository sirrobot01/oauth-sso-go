package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirrobot01/oauth-sso/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserJwtClaims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

type AppJwtClaims struct {
	AppName string `json:"app_name"`
	AppID   string `json:"app_id"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SetCookie(w http.ResponseWriter, name string, value string, expires time.Time) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}

func GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv("SECRET_KEY", "secret")))
}

func GenerateUserToken(userId string, username string, admin bool) (string, error) {
	expirationTime := time.Now().Add(100 * time.Minute)
	claims := UserJwtClaims{
		Username: username,
		Admin:    admin,
		UserID:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oauth2",
			Subject:   userId,
		},
	}
	return GenerateToken(claims)
}

func GenerateAppToken(appId string, appName string) (string, error) {
	expirationTime := time.Now().AddDate(1, 0, 0)
	claims := AppJwtClaims{
		AppName: appName,
		AppID:   appId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oauth2",
			Subject:   appId,
		},
	}
	return GenerateToken(claims)
}

func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}
	if cookie.Value == "" {
		return nil, &ValidationError{
			Message: "invalid password",
		}
	}

	return cookie, nil
}

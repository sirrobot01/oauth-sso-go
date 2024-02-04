package common

import "strings"

var PATHS = make(map[string]map[string]string)

func init() {
	PATHS["user"] = map[string]string{
		"login":    "/user/login/",
		"register": "/user/register/",
	}

	PATHS["app"] = map[string]string{
		"new":  "/app/new/",
		"edit": "/app/{id}/edit/",
		"list": "/app/",
	}

	PATHS[""] = map[string]string{
		"":        "",
		"welcome": "/welcome",
	}

	PATHS["oauth"] = map[string]string{
		"authorize": "/oauth/authorize/",
		"token":     "/oauth/token/",
		"refresh":   "/oauth/refresh/",
		"revoke":    "/oauth/revoke/",
	}
}

func GetPath(name string) string {
	// Split the name into parts user:login
	parts := strings.Split(name, ":")
	return PATHS[parts[0]][parts[1]]
}

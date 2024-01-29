package common

import "strings"

var PATHS = make(map[string]map[string]string)

func init() {
	PATHS["user"] = map[string]string{
		"login":    "/user/login",
		"register": "/user/register",
	}

	PATHS[""] = map[string]string{
		"":        "",
		"welcome": "/welcome",
	}

	PATHS["auth"] = map[string]string{
		"auth":    "/auth/authorize",
		"token":   "/auth/token",
		"refresh": "/auth/refresh",
		"revoke":  "/auth/revoke",
	}
}

func GetPath(name string) string {
	// Split the name into parts user:login
	parts := strings.Split(name, ":")
	return PATHS[parts[0]][parts[1]]
}

package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	bearerToken := headers.Get("Authorization")
	if len(bearerToken) == 0 {
		return "", errors.New("authorization headers doesn't exist")
	}
	list := strings.Split(bearerToken, " ")

	if len(list) != 2 || list[0] != "ApiKey" {
		return "", errors.New("invalid authorization format")
	}

	api_key := list[1]
	return api_key, nil
}

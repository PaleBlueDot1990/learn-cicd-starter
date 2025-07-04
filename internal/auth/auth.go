package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrFooNoAuthHeader = errors.New("no authorization header included")
var ErrFooMalformedAuthHeader = errors.New("malformed authorization header")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrFooNoAuthHeader
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", ErrFooMalformedAuthHeader
	}

	return splitAuth[1], nil
}

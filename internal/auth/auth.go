package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKEY extracts and API Key from the header
//Example:
//Authorisation: APIKey {insert api key here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first path of auth header")
	}
	return vals[1], nil
}

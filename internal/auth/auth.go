package auth

import (
	"errors"
	"net/http"
	"strings"
)

//get api key extracts an api key from headers of http request
//examples
//Authorization: ApiKey {insert api key}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authentication Info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication header found")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authentication header")
	}

	return vals[1], nil
}

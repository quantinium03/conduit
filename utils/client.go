package utils

import (
	"errors"
	"net/http"
)

func GetClientID(r *http.Request) (string, error) {
	clientID := r.Header.Get("CLIENT-ID")
	if clientID == "" {
		return "", errors.New("missing CLIENT-ID in the url parameters")
	}
	return clientID, nil
}

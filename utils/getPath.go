package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func GetPath(r *http.Request) (string, string, error) {
	param := chi.URLParam(r, "path")
	if param == "" {
		return "", "", errors.New("Missing media path")
	}

	pathi, err := base64.RawURLEncoding.DecodeString(param)
	if err != nil {
		return "", "", errors.New("Malformed path. The path should be base64 encoded without padding")
	}

	path := filepath.Clean(string(pathi))
	if !filepath.IsAbs(path) {
		return "", "", errors.New("Malformed path. The path should be an absolute path")
	}

	// TODO: CHECK FOR SOME BAD PATHS

	fileinfo, err := os.Stat(path)
	if err != nil {
		return "", "", err
	}

	hash := sha1.New()
	hash.Write([]byte(path))
	hash.Write([]byte(fileinfo.ModTime().String()))
	sha := hex.EncodeToString(hash.Sum(nil))
	return path, sha, nil
}

package utils

import (
	"log"
	"net/http"
)

func ResponseWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX Error: ", msg)
	}

	type errRes struct {
		Error string `json:"error"`
	}
	ResponseWithJSON(w, code, errRes{
		Error: msg,
	})
}

package handler

import (
	"net/http"

	"github.com/quantinium03/conduit/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
    utils.ResponseWithJSON(w, 200, struct{}{});
}

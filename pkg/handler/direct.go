package handler

import (
	"fmt"
	"net/http"

	"github.com/quantinium03/conduit/utils"
)

func DirectStream(w http.ResponseWriter, r *http.Request) {
	path, _, err := utils.GetPath(r)
	if err != nil {
		utils.ResponseWithErr(w, 500, fmt.Sprintf("Error while directly getting the path: %v", err))
		return
	}
	http.ServeFile(w, r, path)
}

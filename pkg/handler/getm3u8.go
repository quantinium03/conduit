package handler

import (
	"net/http"

)

func Getm3u8(w http.ResponseWriter, r *http.Request) {
	/* id, err := utils.GetClientID(r)
	if err != nil {
		utils.ResponseWithErr(w, 500, err.Error())
		return
	}

	path, sha, err := utils.GetPath(r)
	if err != nil {
		utils.ResponseWithErr(w, 500, err.Error())
		return
	}

	ret, err := transcoder.GetMaster(path, id, sha)
	if err != nil {
		utils.ResponseWithErr(w, 500, err.Error())
	}

	payload := map[string]string{
		"master": ret,
	}

	utils.ResponseWithJSON(w, 200, payload) */
	return
}

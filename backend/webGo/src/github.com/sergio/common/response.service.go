package common

import (
	"encoding/json"
	"net/http"
)

func StatusBadError(w http.ResponseWriter, r *http.Request, ResError string) {
	j, _ := json.Marshal(ResError)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(j)
}

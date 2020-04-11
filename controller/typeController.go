package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jpbmdev/Bodysoft-authentication-ms/services"
)

func getTypesController(w http.ResponseWriter, r *http.Request) {
	types, err := services.GetTypesService()
	if err != "ok" {
		http.Error(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types)
}

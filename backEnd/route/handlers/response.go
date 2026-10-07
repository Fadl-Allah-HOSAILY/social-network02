package handlers

import (
	"backend/model"
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, response model.APIResponse) {

	w.Header().Set("Content-Type", "application/json")
	if response.Code == 0{
        response.Code=200
    }

	w.WriteHeader(response.Code)
    if err := json.NewEncoder(w).Encode(response); err != nil {
		return 
	}
}

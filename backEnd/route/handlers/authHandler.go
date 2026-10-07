package handlers

import (
	"backend/model"
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var userData model.RegisterRequest
	var response model.APIResponse

	if r.Method != http.MethodPost {
		response.Code = http.StatusMethodNotAllowed
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid Method"

		WriteJSONResponse(w, response)

		return
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userData); err != nil {

		response.Code = http.StatusBadRequest
		response.Data = nil
		response.Success = false
		response.PopUp = "Invalid JSON"

		WriteJSONResponse(w, response)

		return
	}
}

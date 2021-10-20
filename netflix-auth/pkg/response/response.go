package response

import (
	"encoding/json"
	"net/http"

	httperror "netflix-auth/pkg/http_error"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, httperror.InternalServerError, http.StatusInternalServerError)
		}
	}
}

func SendErrorResponse(w http.ResponseWriter, err httperror.HTTPError) {
	e, ok := err.(httperror.Wrapper)
	if !ok {
		http.Error(w, httperror.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(e.Code)
	jErr := json.NewEncoder(w).Encode(e)
	if jErr != nil {
		http.Error(w, httperror.InternalServerError, http.StatusInternalServerError)
	}
}

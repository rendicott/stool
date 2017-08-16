package outcome

import (
	"encoding/json"
	"net/http"
)

func OutcomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetOutcomes(a.DB)); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

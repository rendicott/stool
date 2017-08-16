package outcome

import (
	"encoding/json"
	"net/http"

	"github.com/gapi/util"
)

func OutcomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	outcomes, err := GetOutcomes()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(outcomes); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

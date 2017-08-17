package outcome

import (
	"encoding/json"
	"net/http"

	"github.com/gapi/util"
	"strconv"

	"github.com/gorilla/mux"
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

func ShowOutcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	outcomeId, err := strconv.Atoi(vars["outcomeId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	o := Outcome{Id: outcomeId}
	outcome, err := o.GetOutcome()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(outcome); err != nil {
		panic(err)
	}
}

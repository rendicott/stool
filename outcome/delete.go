package outcome

import (
	"net/http"
	"strconv"
	"github.com/gapi/util"

	"github.com/gorilla/mux"
)

func DeleteOutcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	outcomeId, err := strconv.Atoi(vars["outcomeId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	o := Outcome{Id: outcomeId}

	if err := o.DeleteOutcome(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
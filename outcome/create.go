package outcome

import (
	"io/ioutil"
	"io"
	"encoding/json"
	"github.com/gapi/util"
	"net/http"
)

func CreateOutcome(w http.ResponseWriter, r *http.Request) {
	var o Outcome

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if errr := r.Body.Close(); errr != nil {
		panic(errr)
	}
	if errrr := json.Unmarshal(body, &o); errrr != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errrrr := json.NewEncoder(w).Encode(errrr); errrrr != nil {
			panic(errrrr)
		}
	}
	if errrrrr := o.CreateOutcome(); errrrrr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errrrrr.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if errrrrrr := json.NewEncoder(w).Encode(o); errrrrrr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errrrrrr.Error())
		return
	}
}

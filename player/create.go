package player

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gapi/util"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p Player

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if errr := r.Body.Close(); errr != nil {
		panic(errr)
	}
	if errrr := json.Unmarshal(body, &p); errrr != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errrrr := json.NewEncoder(w).Encode(errrr); errrrr != nil {
			panic(errrrr)
		}
	}
	if errrrrr := p.CreatePlayer(); errrrrr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errrrrr.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if errrrrrr := json.NewEncoder(w).Encode(p); errrrrrr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errrrrrr.Error())
		return
	}
}

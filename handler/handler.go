package handler

import (
	"github.com/ZW-Rays/golang-login/utils"
	"net/http"
)

type ServiceFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

func MakeHandler(service ServiceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			return
		}

		response, err := service(w, r)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

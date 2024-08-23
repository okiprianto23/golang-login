package handler

import (
	"encoding/json"
	"errors"
	"github.com/ZW-Rays/golang-login/model"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	//check method http
	if r.Method != http.MethodPost {
		return nil, errors.New("method is not post")
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid request payload")
	}

	if user.Username == "admin" && user.Password == "admin" {
		return map[string]string{"message": "login successful - admin"}, nil
	} else if user.Username == "user" && user.Password == "user" {
		return map[string]string{"message": "login successful - user"}, nil
	}

	return nil, errors.New("authentication failed")
}

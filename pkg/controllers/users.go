package controllers

import (
	"dashboard-api/pkg/dbms"
	"dashboard-api/pkg/dbms/model"
	"dashboard-api/pkg/handlers/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//Login irá controlar a autenticação de um usuário
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, errors.New("parâmetros inválidos"))
		return
	}

	db, err := dbms.Connect()
	defer db.Close()

	responses.JSON(w, http.StatusOK, user)
}

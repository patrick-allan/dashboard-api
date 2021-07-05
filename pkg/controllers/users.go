package controllers

import (
	"dashboard-api/pkg/dbms"
	"dashboard-api/pkg/dbms/model"
	"dashboard-api/pkg/dbms/repository"
	"dashboard-api/pkg/handlers/responses"
	"dashboard-api/pkg/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userDB, err := repo.Login(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if (model.User{}) == userDB {
		responses.Err(w, http.StatusInternalServerError, errors.New("e-mail inválido"))
		return
	}

	if err = security.VerificarSenha(userDB.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, errors.New("senha incorreta"))
		return
	}

	token, err := security.MakeToken(userDB.Email, strconv.Itoa(userDB.Id), userDB.Name)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

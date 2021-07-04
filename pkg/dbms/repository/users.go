package repository

import (
	"dashboard-api/pkg/dbms/model"
	"database/sql"
	"errors"
)

//Users representa o repositório de Usuários
type Users struct {
	db *sql.DB
}

//NewUsersRepository inicializa o repositorio de Usuários
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

//Login irá retornar as informações de login de um usuário
func (repository Users) Login(email string) (model.User, error) {
	sql, err := repository.db.Query("select id, name, email, password from users where email = ?", email)
	if err != nil {
		return model.User{}, errors.New("error sql prepare - " + err.Error())
	}
	defer sql.Close()

	var user model.User
	if sql.Next() {
		if err = sql.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

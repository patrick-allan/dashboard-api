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
		return model.User{}, errors.New("error sql login - " + err.Error())
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

//Register irá incluir um novo usuário
func (repository Users) Register(user model.User) (int, error) {
	sql, err := repository.db.Prepare(
		"insert into users (name, email, username, password, birthdate, document, created_at) values (?,?,?,?,?,?,?)",
	)
	if err != nil {
		return 0, errors.New("error sql register user - " + err.Error())
	}
	defer sql.Close()

	result, err := sql.Exec(
		user.Name,
		user.Email,
		user.Username,
		user.Password,
		user.Birthdate,
		user.Document,
		user.Created_at)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("não foi possível obter o usuário inserido")
	}

	return int(lastId), nil
}

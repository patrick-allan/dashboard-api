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
func (repository Users) Login(email, password string) (model.User, error) {
	return model.User{}, errors.New("repositorio de usuário, função login deverá ser implementada")
}

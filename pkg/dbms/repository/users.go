package repository

import (
	"dashboard-api/pkg/dbms/model"
	"database/sql"
)

//Users representa o repositório de Usuários
type Users struct {
	db *sql.DB
}

//NewUsersRepository inicializa o repositorio de Usuários
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Login(email, password string) (model.User, error) {

}

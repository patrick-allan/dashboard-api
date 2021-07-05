package dbms

import (
	"dashboard-api/pkg/config"
	"database/sql"
	"errors"
	"fmt"

	//sqllite
	_ "github.com/mattn/go-sqlite3"
)

//Connect retorna *sql.DB conectar no DB
func Connect() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("./%s?cache=shared&mode=memory",
		config.DB_DATABASE)

	//Initialize connection object.
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, errors.New("DB Connect error - " + err.Error())
	}
	//Executa um ping no DB
	err = db.Ping()
	if err != nil {
		return nil, errors.New("DB Ping error - " + err.Error())
	}

	return db, nil
}

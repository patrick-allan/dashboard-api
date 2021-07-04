package dbms

import (
	"dashboard-api/pkg/config"
	"database/sql"
	"errors"
	"fmt"

	//sqllite
	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//Connect retorna *sql.DB conectar no DB
func Connect() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("file:%s?cache=shared&mode=memory",
		config.DB_DATABASE)

	// Initialize connection object.
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, errors.New("DB Connect error - " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("DB Ping error - " + err.Error())
	}

	return db, nil
}

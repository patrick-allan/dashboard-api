package config

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//DB_HOST é nome/ip do DB
	DB_HOST = ""
	//DB_HOST é nome/ip do DB Replicado
	DB_HOST_REP = ""
	//DB_PORT é a porta de conexão do DB
	DB_PORT = 0
	//DB_DATABASE é o nome do banco de dados no DB
	DB_DATABASE = ""
	//DB_USERNAME é o usuário de conexão no DB
	DB_USERNAME = ""
	//DB_PASSWORD é a senha de conexão no DB
	DB_PASSWORD = ""
	//APP_PORT é a porta onde a API vai estar executando
	APP_PORT = ""
	//APP_KEY é a chave de assinatura do JWT-Token
	APP_KEY = ""
	//JWT_TTL JWT time to live - tempo (minutos) em que o token é valido
	JWT_TTL = 0
)

//Load vai inicializar as variaveis de ambiente
func Load() {
	var err error
	var gravaEnv bool

	if err = godotenv.Load(); err != nil {
		log.Fatal(errors.New("não foi possível carregar o arquivo .env"))
	}

	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("não foi possível ler o arquivo .env")
	}

	DB_HOST = myEnv["DB_HOST"]
	DB_HOST_REP = myEnv["DB_HOST_REP"]
	DB_PORT, err = strconv.Atoi(myEnv["DB_PORT"])
	if err != nil {
		log.Fatal(err)
	}
	DB_DATABASE = myEnv["DB_DATABASE"]
	DB_USERNAME = myEnv["DB_USERNAME"]
	DB_PASSWORD = myEnv["DB_PASSWORD"]
	APP_PORT = myEnv["APP_PORT"]
	APP_KEY = myEnv["APP_KEY"]

	if APP_KEY == "" {
		APP_KEY, err = generateAppKey()
		if err != nil {
			log.Fatal(err)
		}
		myEnv["APP_KEY"] = APP_KEY
		gravaEnv = true
	}
	JWT_TTL, err = strconv.Atoi(myEnv["JWT_TTL"])
	if err != nil {
		log.Fatal(err)
	}
	if JWT_TTL == 0 {
		JWT_TTL = 360
		myEnv["JWT_TTL"] = strconv.Itoa(JWT_TTL)
		gravaEnv = true
	}

	if gravaEnv {
		if err = godotenv.Write(myEnv, "./.env"); err != nil {
			log.Fatal("não foi possível gravar o arquivo .env")
		}
	}
}

//generateAppKey gera a chave de assinatura do token APP_KEY environment
func generateAppKey() (string, error) {
	chave := make([]byte, 64)
	if _, err := rand.Read(chave); err != nil {
		return "", err
	}
	strBase64 := base64.StdEncoding.EncodeToString(chave)
	return strBase64, nil
}

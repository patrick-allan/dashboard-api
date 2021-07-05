package security

import (
	"dashboard-api/pkg/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Login representa o retorno de um login
type Login struct {
	Token  string `json:"token,omitempty"`
	Expire string `json:"expire,omitempty"`
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
}

//MakeToken retorna um token assinado com as permissoes do usuário
func MakeToken(userEmail, userId, userName string) (Login, error) {
	jwt_ttl := time.Now().Add(time.Minute * time.Duration(config.JWT_TTL)).Unix()
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = jwt_ttl
	permissoes["id"] = userId
	permissoes["email"] = userEmail

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	var login Login
	var err error
	login.Token, err = token.SignedString([]byte(config.APP_KEY))
	if err != nil {
		return Login{}, err
	}
	login.Expire = strconv.Itoa(int(jwt_ttl))
	login.Name = userName
	login.Id, err = strconv.Atoi(userId)
	if err != nil {
		return Login{}, err
	}
	return login, nil
}

//ExtracUserEmail retorna o email contigo no Token
func ExtracUserEmail(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnAppKey)
	if err != nil {
		return "", err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail := fmt.Sprintf("%v", permissoes["email"])
		return userEmail, nil
	}

	return "", errors.New("token inválido")
}

//ValidateToken verifica que o token da requisição é válido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnAppKey)
	if err != nil {
		return errors.New(err.Error())
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

//extractToken retorna o token que está no HEADER-Authorization
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

//returnAppKey retorna []Byte APP_Key validando o método de assinatura do token
func returnAppKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return []byte(config.APP_KEY), nil
}

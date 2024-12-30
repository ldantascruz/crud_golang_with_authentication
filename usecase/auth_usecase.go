package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase struct {
}

func NewAuthUsecase() AuthUsecase {
	return AuthUsecase{}
}

var secretKey = os.Getenv("SECRET_KEY")

func (au *AuthUsecase) Login(username, password string) (string, error) {
	// Aqui você deve implementar a validação do usuário. Para testes:
	if username != "admin" || password != "123456" {
		return "", jwt.ErrSignatureInvalid
	}

	// Gerar o token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	// Assinar o token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package jwt


import jwt "github.com/dgrijalva/jwt-go"


func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	return _, nil
}
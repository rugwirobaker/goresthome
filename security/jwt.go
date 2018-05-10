package security

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generates a rsa 256 signed JWT token.
func GenerateJWT(name, role string, signKey interface{}) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set claims for JWT token
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}
	//set the expire time for JWT token
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	t.Claims = claims

	//sign the token
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	//log.Println(tokenString)
	return tokenString, nil
}

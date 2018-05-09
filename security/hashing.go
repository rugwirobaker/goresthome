package security

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword returns a bcrypt hashed password.
//It wraps GenerateFromPassword(password []byte, cost int) ([]byte, error)
func HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 14)
	return hash, err
}

//CheckPasswordHash verifies whether a given password , matches a gives
// a given Hash
func CheckPasswordHash(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

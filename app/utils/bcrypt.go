package utils

import "golang.org/x/crypto/bcrypt"

func PasswordHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

func PasswordVerify(pwd string, hash string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false
	} else {
		return true
	}
}

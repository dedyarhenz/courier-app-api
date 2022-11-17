package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != err {
		return "", err

	}

	return string(hash), nil
}

func ComparePassword(password string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input))
	if err != nil {
		return false
	}

	return true
}

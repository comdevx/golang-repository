package helper

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Password(password string) (string, error) {

	password = password + os.Getenv("PASSWORD_SECRET")
	encode, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encode), nil
}

func ComparePassword(password1, password2 string) error {

	password1 = password1 + os.Getenv("PASSWORD_SECRET")
	password2 = password2 + os.Getenv("PASSWORD_SECRET")
	if err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2)); err != nil {
		return err
	}

	return nil
}

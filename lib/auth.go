package lib

import (
	"errors"
	"CMP1066/models"
	"crypto/sha256"
)

func Authenticate(nick string, password string) (user *models.User, err error) {
	msg := "invalid email or password."
	user = &models.User{Nick: nick, Status: true}

	if err := user.Read("Nick","Status"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.Id < 1 {
		// No user
		return user, errors.New(msg)
	} else if user.Password != Crypto(password) {
		// No matched password
		return user, errors.New(msg)
	} else {
		return user, nil
	}
}

func Crypto(p string) string {
	hashSenha := sha256.Sum256([]byte(p))
	return string(hashSenha[:])
}
